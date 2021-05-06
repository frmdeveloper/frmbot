const {
   WAConnection,
   MessageType,
   Presence,
   MessageOptions,
   Mimetype,
   WALocationMessage,
   WA_MESSAGE_STUB_TYPES,
   ReconnectMode,
   ProxyAgent,
   GroupSettingChange,
   ChatModification,
   waChatKey,
   mentionedJid,
   processTime,
   WA_DEFAULT_EPHEMERAL
} = require("@adiwajshing/baileys")

exports.client = client = new WAConnection()
client.logger.level = 'warn'
console.log(banner.string)