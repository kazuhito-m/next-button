# 手動操作

fabricのプロビジョニングだけでは片付けられなかった、手動操作を書き残す。

## SDカードの拡張

最初は、何故か小さい状態から始まる。

その領域を目一杯広げる操作を行う。

備え付けのツールにて設定

RaspPIにログインし…

```bash
sudo raspi-config
```

から`1 Expand Filesystem`を指定、locate,timezoneを変更する。

## locale,timezone変更

備え付けのツールにて設定

RaspPIにログインし…

```bash
sudo raspi-config
```

から`4 Internationalisation Options`を指定、locate,timezoneを変更する。
