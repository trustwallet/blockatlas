# -*- mode: Python -*-

local_resource(
  'lint+tests',
  'make go-lint && make go-test && go-integration',
  trigger_mode=TRIGGER_MODE_MANUAL, auto_init=False
)

docker_build("trust/blockatlas:api-local", ".", build_args={"SERVICE":"api"})
docker_build("trust/blockatlas:parser-local", ".", build_args={"SERVICE":"parser"})
docker_build("trust/blockatlas:consumer-local", ".", build_args={"SERVICE":"consumer"})

yaml = helm(
  'deployment/charts/blockatlas',
  name='local',
  namespace='tilt-blockatlas-local',
  values=['./deployment/charts/blockatlas/values.local.yaml']
)

# k8s namespace bootstrap
local('kubectl create namespace tilt-blockatlas-local || echo 1')

k8s_yaml(yaml)
k8s_resource('api', port_forwards=8420)

k8s_resource('postgres', port_forwards=8586)
k8s_resource('rabbitmq', port_forwards='9596:15672')
