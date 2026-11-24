# bantuan/

Padanan wa33/bantuan/: aset statis (font .ttf, stiker .webp, audio, cookie,
berkas workflow .yml) + subfolder mywagua.

- Aset (font/stiker/dll) TIDAK disalin otomatis oleh generator skeleton —
  salin manual dari wa33/bantuan/ pas fitur yang memakainya diporting.
- bantuan/mywagua SENGAJA tidak diporting: itu binding JS<->Go buat
  whatsmeow. Di gow1, whatsmeow/hypermeow dipakai langsung dari Go (conn.go),
  jadi lapis binding-nya hilang — memang itu tujuan porting full Go.
