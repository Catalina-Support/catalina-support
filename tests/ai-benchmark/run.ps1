# Banco de pruebas de modelos locales para Catalina Support.
#
# No toca el motor activo ni producción: descarga los GGUF en un volumen aparte, levanta un
# llama-server efímero por modelo en el puerto 18080 y le manda exactamente los mismos prompts que
# arma la aplicación. Los resultados quedan en tests/ai-benchmark/results/.
#
# Uso:
#   pwsh -File tests/ai-benchmark/run.ps1
#   pwsh -File tests/ai-benchmark/run.ps1 -Only qwen3-0.6b,qwen3-1.7b
#
# Requiere Docker. La primera pasada descarga unos 3,5 GB de modelos.

param(
  [string[]]$Only = @(),
  [string]$Image = 'ghcr.io/ggml-org/llama.cpp:server-b11206',
  [string]$Volume = 'catalina_support_ai_bench_models',
  [int]$Port = 18080,
  [int]$Threads = 2,
  [switch]$ThinkingOff,
  [string]$Suffix = ''
)

$ErrorActionPreference = 'Stop'
$root = (Resolve-Path (Join-Path $PSScriptRoot '../..')).Path
$bench = Join-Path $root 'tests/ai-benchmark'
$results = Join-Path $bench 'results'
New-Item -ItemType Directory -Force -Path $results | Out-Null

$modelos = [ordered]@{
  'qwen2.5-1.5b' = @{
    archivo = 'qwen2.5-1.5b-instruct-q4_k_m.gguf'
    url = 'https://huggingface.co/Qwen/Qwen2.5-1.5B-Instruct-GGUF/resolve/main/qwen2.5-1.5b-instruct-q4_k_m.gguf?download=true'
    nota = 'modelo actual'
  }
  'qwen3-0.6b' = @{
    archivo = 'Qwen_Qwen3-0.6B-Q4_K_M.gguf'
    url = 'https://huggingface.co/bartowski/Qwen_Qwen3-0.6B-GGUF/resolve/main/Qwen_Qwen3-0.6B-Q4_K_M.gguf?download=true'
    nota = 'candidato más rápido'
  }
  'qwen3-1.7b' = @{
    archivo = 'Qwen_Qwen3-1.7B-Q4_K_M.gguf'
    url = 'https://huggingface.co/bartowski/Qwen_Qwen3-1.7B-GGUF/resolve/main/Qwen_Qwen3-1.7B-Q4_K_M.gguf?download=true'
    nota = 'candidato equilibrado'
  }
  'gemma3-1b' = @{
    archivo = 'google_gemma-3-1b-it-Q4_K_M.gguf'
    url = 'https://huggingface.co/bartowski/google_gemma-3-1b-it-GGUF/resolve/main/google_gemma-3-1b-it-Q4_K_M.gguf?download=true'
    nota = 'candidato multilingüe'
  }
  'qwen2.5-0.5b' = @{
    archivo = 'qwen2.5-0.5b-instruct-q4_k_m.gguf'
    url = 'https://huggingface.co/Qwen/Qwen2.5-0.5B-Instruct-GGUF/resolve/main/qwen2.5-0.5b-instruct-q4_k_m.gguf?download=true'
    nota = 'alternativa mínima sin razonamiento'
  }
}

$seleccionados = if ($Only.Count -gt 0) { $Only } else { $modelos.Keys }
foreach ($nombre in $seleccionados) {
  if (-not $modelos.Contains($nombre)) {
    throw "Modelo desconocido: $nombre. Usa: $($modelos.Keys -join ', ')"
  }
}

Write-Host '[bench] Preparando imagen y volumen' -ForegroundColor Cyan
docker pull $Image | Out-Host
docker volume create $Volume | Out-Null

foreach ($nombre in $seleccionados) {
  $modelo = $modelos[$nombre]
  Write-Host "[bench] Descargando $nombre ($($modelo.nota))" -ForegroundColor Cyan
  $script = "if [ -s /models/$($modelo.archivo) ]; then echo 'ya existe'; else wget -O /models/$($modelo.archivo) '$($modelo.url)'; fi"
  docker run --rm -v "${Volume}:/models" alpine sh -lc $script | Out-Host
}

foreach ($nombre in $seleccionados) {
  $modelo = $modelos[$nombre]
  $contenedor = "ai-bench-$($nombre -replace '[^a-zA-Z0-9]', '-')"
  $archivo = "/models/$($modelo.archivo)"

  Write-Host "[bench] Arrancando $nombre" -ForegroundColor Cyan
  docker rm -f $contenedor 2>$null | Out-Null
  $extra = @()
  if ($ThinkingOff -and $nombre -like 'qwen3-*') {
    $extra = @('--reasoning', 'off')
  }
  docker run -d --rm --name $contenedor -p "127.0.0.1:${Port}:8080" -v "${Volume}:/models" $Image `
    --model $archivo --alias $nombre --host 0.0.0.0 --port 8080 --ctx-size 4096 --threads $Threads `
    --parallel 1 --n-predict 512 --no-webui @extra | Out-Null

  try {
    $listo = $false
    for ($i = 0; $i -lt 180; $i++) {
      Start-Sleep -Seconds 1
      try {
        $salud = Invoke-RestMethod -Uri "http://127.0.0.1:${Port}/health" -TimeoutSec 2
        if ($salud.status -eq 'ok' -or $salud -eq 'ok') { $listo = $true; break }
      } catch { }
    }
    if (-not $listo) { throw "El modelo $nombre no alcanzó salud" }

    Write-Host "[bench] Evaluando $nombre" -ForegroundColor Cyan
    $salida = "/bench/results/$nombre$Suffix.json"
    docker run --rm --add-host=host.docker.internal:host-gateway `
      -v "${bench}:/bench" python:3.12-slim `
      python /bench/run.py --base-url http://host.docker.internal:${Port} --model $nombre `
      --cases /bench/cases.json --out $salida | Out-Host
  }
  finally {
    docker stop $contenedor 2>$null | Out-Null
  }
}

Write-Host "[bench] Resultados en $results" -ForegroundColor Green
