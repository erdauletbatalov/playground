-- Открываем порт для доступа по iproto
box.cfg({listen="127.0.0.1:3301"})
-- Создаём пользователя для подключения
box.schema.user.create('storage', {password='passw0rd', if_not_exists=true})
-- Даём все-все права
box.schema.user.grant('storage', 'super', nil, nil, {if_not_exists=true})
-- Чуть настраиваем сериализатор в iproto, чтобы не ругался на несериализуемые типы
require('msgpack').cfg{encode_invalid_as_nil = true}