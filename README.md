# pcba-test
This is automated testing of the PCBAs for the cacophony project

## Setup
This is intended to be tested on a Raspberry Pi 3 B
- Install thermal-recorder https://github.com/TheCacophonyProject/thermal-recorder/releases
- Enable thermal-recorder service `sudo systemctl enable thermal-recorder`
- Enable leptond service `sudo sudo systemctl enable leptond`
- Replace `/boot/config.txt` on Raspberry Pi with `config.txt`
- Setup wifi on Raspberry Pi so you can connect to it on over the same network. https://www.raspberrypi.org/documentation/configuration/wireless/wireless-cli.md
- Downoad and install latest pcba-test package.
- Have a ATtiny85 programmed (at 8MHz) with the latest code from https://github.com/TheCacophonyProject/attiny
- (optional) change hostname of the Raspberry Pi if you are setting up more than one test device

## Testing
- Power off Raspberry Pi
- Put Hat on Raspberry Pi
- Attach:
  - RTC battery
  - Thermal camera
  - Speaker
  - ATtiny
  - Reset/LED button
  - Attach USB cable from Raspberry Pi to HAT
  - Put USB drive in HAT
- Power up through power plug on HAT
- Open a web browser on a device that is on the same wii as the Raspberry Pi and go to `<RPi-devicename>.local`
- Press `Run Tests` and wait for it to finish
- Press `Play Sound` and check that it played a sound
- Press `View Camera` and check the output from the camera (this can take a minute or two sometimes to start from boot)
