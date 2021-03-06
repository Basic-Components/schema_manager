"""sanic默认的app构造组件.

组件名字会被改为项目名.

注册模块组件使用`xxx.init_app(app)`;

使用数据库组件放在hooks模块中在启动项目时挂载
"""
from sanic import Sanic
from api import restapi
from hooks import hooks
from exception import excep
from const import SERVICE_NAME
from log import (
    LOGGING_CONFIG_JSON,
    set_mail_log
)


def init_app(config):
    log_config = None
    if config["SET_LOG_FMT"] == "json":
        log_config = LOGGING_CONFIG_JSON
    app = Sanic(SERVICE_NAME, log_config=log_config)
    app.config.update(
        config
    )
    if app.config.SET_LOG_MAIL_LOG is True:
        set_mail_log(app)
    restapi.init_app(app)
    hooks.init_app(app)
    excep.init_app(app)
    return app