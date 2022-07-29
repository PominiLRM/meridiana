
<img src="public/img/meridiana.png">  

The open-source platform for monitoring and observability

[![License](https://img.shields.io/github/license/grafana/grafana)](LICENSE)
[![Drone](https://drone.grafana.net/api/badges/grafana/grafana/status.svg)](https://drone.grafana.net/grafana/grafana)
[![Go Report Card](https://goreportcard.com/badge/github.com/grafana/grafana)](https://goreportcard.com/report/github.com/grafana/grafana)

Grafana allows you to query, visualize, alert on and understand your metrics no matter where they are stored. Create, explore, and share dashboards with your team and foster a data-driven culture:

- **Visualizations:** Fast and flexible client side graphs with a multitude of options. Panel plugins offer many different ways to visualize metrics and logs.
- **Dynamic Dashboards:** Create dynamic & reusable dashboards with template variables that appear as dropdowns at the top of the dashboard.
- **Explore Metrics:** Explore your data through ad-hoc queries and dynamic drilldown. Split view and compare different time ranges, queries and data sources side by side.
- **Explore Logs:** Experience the magic of switching from metrics to logs with preserved label filters. Quickly search through all your logs or streaming them live.
- **Alerting:** Visually define alert rules for your most important metrics. Grafana will continuously evaluate and send notifications to systems like Slack, PagerDuty, VictorOps, OpsGenie.
- **Mixed Data Sources:** Mix different data sources in the same graph! You can specify a data source on a per-query basis. This works for even custom datasources.

## Documentation

The Grafana documentation is available at [grafana.com/docs](https://grafana.com/docs/).

## Developer environment

SO: Ubuntu  

### GO

sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go

go --version

### Nodejs LTS
curl -fsSL https://deb.nodesource.com/setup_lts.x | sudo -E bash -
sudo apt-get install -y nodejs

node --version

### optional
apt install npm
npm install -g yarn

yarn dlx @yarnpkg/sdks vscode

### Yarn
apt-get install yarn

yarn --version

### git
apt-get install git

git --version


### grafana sources
sudo su

yarn install --immutable
yarn start

sudo apt install make

make run

ulimit -S -n 4096


### Docker

apt install docker.io

docker --version

systemctl start docker

systemctl enable docker

### Info

https://grafana.com/docs/grafana/latest/installation/docker/?msclkid=7ff4adaec54911ec80ec13ed85010633


## License

Meridiana is distributed under [AGPL-3.0-only](LICENSE).
