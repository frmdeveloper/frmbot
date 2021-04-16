const help = (pushname, prefix, botName, ownerName, reqXp, getLevelingLevel, sender, _registered, uangku) => {
        return `🔰 -----[ *MENU ${botName}* ]----- 🔰
Hallo, ${pushname} 👋
Semoga harimu Menyenangkan User, ${sender.split("@")[0]}
┏━━━━━━━━━━━━━━━━━┓
┃╭───────────────
┃│➸ NAMA : ${pushname}
┃│➸ UANG : Rp:${uangku}
┃│➸ XP : ${reqXp}
┃│➸ LEVEL : ${getLevelingLevel(sender)}
┃│➸ USER ${botName} : ${_registered.length}
┃╰───────────────
┗━━━━━━━━━━━━━━━━━┛
Berikut adalah fitur yang ada pada bot ini!✨
Jika tidak paham ketik *${prefix}bingungcok*
┏━━━━━━━━━━━━━━━━━┓
┃╭───────────────
┃│➸ *${prefix}info*
┃│➸ *${prefix}donasi*
┃│➸ *${prefix}owner*
┃│───────────────
┃│➸ *${prefix}simplemenu2*
┃│➸ *${prefix}makermenu2*
┃│➸ *${prefix}gabutmenu2*
┃│➸ *${prefix}downloadmenu2*
┃│➸ *${prefix}randommenu2*
┃│➸ *${prefix}dompetmenu2*
┃│➸ *${prefix}othermenu2*
┃│➸ *${prefix}groupmenu2*
┃│➸ *${prefix}soundmenu2*
┃│➸ *${prefix}ownermenu2*
┃╰───────────────
┗━━━━━━━━━━━━━━━━━┛
🔰 -----[ *POWERED BY ${ownerName}* ]----- 🔰`
}
exports.help = help
