all:
	yarn install --pure-lockfile
	yarn build

dev:
	yarn dev

watch:
	yarn watch

test:
	yarn test

test-watch:
	yarn test-watch

add-sync:
	# install syncyarnlock globally
	yarn global add syncyarnlock # install syncyarnlock globally

upgrade-packages: add-sync
	# update dependencies, updates yarn.lock
	yarn upgrade
	# updates package.json with versions installed from yarn.lock
	syncyarnlock -g -l -s -a "github:CorpGlory/types-grafana"
	# updates yarn.lock with current version constraint from package.json
	yarn install

