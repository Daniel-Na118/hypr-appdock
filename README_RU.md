# hypr-appdock
### ?н?е?ак?ивна? док-панел? дл? Hyprland

<img width="1360" height="768" alt="250725_16h02m52s_screenshot" src="https://github.com/user-attachments/assets/041d2cf6-13ba-4c89-a960-1903073ff2d4" />
<img width="1360" height="768" alt="250725_16h03m09s_screenshot" src="https://github.com/user-attachments/assets/0c1ad8ca-37c1-4fd6-a48d-46f74c2d2609" />

[![YouTube](https://img.shields.io/badge/YouTube-?идео-FF0000?logo=youtube)](https://youtu.be/HHUZWHfNAl0?si=ZrRv2ggnPBEBS5oY)
[![AUR](https://img.shields.io/badge/AUR-Package-1793D1?logo=arch-linux)](https://aur.archlinux.org/packages/hypr-appdock)

## У??ановка

### ?ави?имо??и

- `go` (make)
- `gtk3`
- `gtk-layer-shell`

### У??ановка
! ?е?ва? ?бо?ка може? занима?? к?айне много в?емени из за п?ив?зки gtk3 !
```bash
git clone https://github.com/lotos-linux/hypr-appdock.git
cd hypr-appdock
make get
make build
make install
```

### Удаление
```bash
make uninstall
```

### ?окал?н?й зап??к (dev mode)
```bash
make exec
```

## ?ап??к

### ?а?аме??? зап??ка:

```text
  -config string
    	config file (default "~/.config/hypr-appdock")
  -dev
    	enable developer mode
  -log-level string
    	log level (default "info")
  -theme string
    	theme dir
```
#### ??е па?аме??? ?вл????? необ?за?ел?н?ми.

?он?иг??а?и? и ?ем? по ?мол?ани? ??ав????? в `/etc/hypr-appdock`
??и пе?вом зап??ке копи?????? в `~/.config/hypr-appdock`
### ?обав??е зап??к в `hyprland.conf`:

```text
exec-once = hypr-appdock
bind = Super, D, exec, hypr-appdock
```

### ? на???ой?е бл?? е?ли он вам н?жен
```text
layerrule = blur true,match:namespace hypr-appdock
layerrule = ignore_alpha 0,match:namespace hypr-appdock
layerrule = blur true,match:namespace dock-popup
layerrule = ignore_alpha 0,match:namespace dock-popup
```

#### ?ок подде?живае? ?ол?ко один зап??енн?й ?кземпл??, ?ак ??о пов?о?н?й зап??к зак?ое? п?ед?д???ий.

## ?а???ойка

#### !?ажно! 
??ли ? ва? ?же ??о?л `hypr-appdock` или в? обновл?е?е?? ?о п?о?? об?а?и?? внимание ??о на?ина? ? ве??ии 1.2.0 ?епе?? и?пол?з?е??? `ini` подобн?й ?ома? кон?иг??а?ии. ? е?ли ? ва? в пол?зова?ел??кой папке о??али?? ??а??е кон?иги об?за?ел?но ?дали?е и? или пе?еме??и?е. ?а?ем п?о??о адап?и??й?е ва?и на???ойки под нов?й ?о?ма?

### ? `hypr-appdock.conf` до???пн? ?акие па?аме???

```ini
[General]
CurrentTheme = lotos

# Icon size (px) (default 23)
IconSize = 23

# Window overlay layer height (background, bottom, top, overlay) (default top)
Layer = top

# Exclusive Zone (true, false) (default true)
Exclusive = true

# SmartView (true, false) (default false)
SmartView = false

# Window position on screen (top, bottom, left, right) (default bottom)
Position = bottom

# Delay before hiding the dock (ms) (default 400)
AutoHideDelay = 400   # Only for SmartView

# Use system gap (true, false) (default true)
SystemGapUsed = true

# Indent from the edge of the screen (px) (default 8)
Margin = 8

# Distance of the context menu from the window (px) (default 5)
ContextPos = 5

[General.preview]
# Window thumbnail mode selection (none, live, static) (default none)
Mode = none
# "none"   - disabled (text menus)
# "static" - last window frame
# "live"   - window streaming
      
# !WARNING! 
# BY SETTING "Mode" TO "live" OR "static", YOU AGREE TO THE CAPTURE 
# OF WINDOW CONTENTS.
# THE "hypr-appdock" PROGRAM DOES NOT COLLECT, STORE, OR TRANSMIT ANY DATA.
# WINDOW CAPTURE OCCURS ONLY FOR THE DURATION OF THE THUMBNAIL DISPLAY!
#   
# Source code: https://github.com/lotos-linux/hypr-appdock

# Live preview fps (0 - ?? (default 30)
FPS = 30

# Live preview bufferSize (1 - 20) (default 5)
BufferSize = 5

# Popup show/hide/move delays (ms)
ShowDelay = 500  # (default 500)
HideDelay = 350  # (default 350)
MoveDelay = 100  # (default 100)
```
#### ??ли па?аме?? не ?казан зна?ение б?де? в???авлено по ?мол?ани?

## ?азбе?ем нео?евидн?е па?аме???

### SmartView
Ч?о ?о на подобии ав?о?к???и?, е?ли `true` ?о док на?оди???? под в?еми окнами, но е?ли ?ве??и к???о? м??и к к?а? ?к?ана - док поднимае??? над ними

### Exclusive
?к?иви??е? о?обое поведение ?ло? п?и ко?о?ом ?айлингов?е окна не пе?ек??ва?? док

### SystemGapUsed
- ??и `SystemGapUsed = true` док б?де? задава?? дл? ?еб? о????п о? к?а? ?к?ана бе?? зна?ение из кон?иг??а?ии `hyprland`, а конк?е?но зна?ени? `general:gaps_out`, п?и ??ом док динами?е?ки б?де? под?ва??ва?? изменение кон?иг??а?ии `hyprland`
- ??и `SystemGapUsed = false` о????п о? к?а? ?к?ана б?де? задава???? па?аме??ом `Margin`

### General.preview
- `ShowDelay`, `HideDelay`, `MoveDelay` - заде?жки дей??вий попапа п?ев?? в мили?ек?нда?
- `FPS`, `BufferSize` - и?пол?з????? ?ол?ко п?и `Mode = live`


#### ?а???ойки вне?него вида п?ев?? п?ои??оди? ?е?ез ?айл? ?ем?



### ?аклепленн?е п?иложени? ??ан????? в ?айле `~/.local/share/hypr-appdock/pinned`
?л? зак?еплени? о?к?ой?е кон?е??ное мен? п?иложени? в доке и нажмине `pin`/`unpin`
#### ?ап?име?
```text
firefox
code-oss
kitty
org.telegram.desktop
nemo
org.kde.ark
sublime_text
qt6ct
one.ablaze.floorp
```
?? може?е мен??? его в ???н??. ?о за?ем? ¯\_(??_/¯

## Тем?

#### Тем? на?од????? в папке `~/.config/hypr-appdock/themes/`

### Тема ?о??ои? из
- `theme.conf`
- `style.css`
- ?апка ? `svg` ?айлами дл? индика?ии коли?е??ва зап??енн?? п?иложени? (?мо??и?е [themes_RU.md](https://github.com/lotos-linux/hypr-appdock/blob/main/docs/customize/themes_RU.md))

### ?он?иг ?ем?
```ini
[Theme]
# Distance between elements (px) (default 9)
Spacing = 5


[Theme.preview]
# Size (px) (default 120)
Size = 120

# Image/Stream border-radius (px) (default 0)
BorderRadius = 0

# Popup padding (px) (default 10)
Padding = 10
```
#### Файл `style.css` к???и?е как ?о?и?е. ?озже ?дела? под?обн?? док?мен?а?и? по ??илиза?ии
