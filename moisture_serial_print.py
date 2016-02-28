#!/usr/bin/env python
import serial

ser = serial.Serial(
    port='/dev/ttyUSB0',
    baudrate = 9600,
    parity=serial.PARITY_NONE,
    stopbits=serial.STOPBITS_ONE,
    bytesize=serial.EIGHTBITS,
    timeout=1
)
ser.write('x')

readStr = ""

while 1:
    readStr += ser.readline()
    if readStr.find("EOF") != -1:
        break

print(readStr)
