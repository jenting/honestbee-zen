#!/usr/bin/env python

import yaml
import os

deploy_env = os.environ.get('ENV', 'staging')

if deploy_env == "production" :
    print ("deploy_env is production , set env as PROD")
    env = 'PROD'
else:
    print ("deploy_env is not production , set env as STAGING")
    env = 'STAGING'


with open("env.yml", 'r') as f:
    config = yaml.load(f)
    config['http']['listen_addr'] = os.environ.get('PORT')
    config['database']['host'] = os.environ.get(env + "_" + 'DB_HOST')
    config['database']['db_name'] = os.environ.get(env + "_" + 'DB_NAME')
    config['database']['user'] = os.environ.get(env + "_" + 'DB_USER')
    config['database']['password'] = os.environ.get(env + "_" + 'DB_PASSWORD')
    config['cache']['host'] = os.environ.get(env + "_" + 'REDIS_HOST')
    config['zendesk']['auth_token'] = os.environ.get(env + "_" + 'ZENDESK_AUTH_TOKEN')
    config['datadog']['env'] = os.environ.get(env + "_" + 'DATADOG_ENV')

with open("env.yml", 'w') as f:
    yaml.dump(config, f, default_flow_style=False)


print ("env.yml setup completed ")
