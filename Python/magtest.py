import smbus
from time import sleep
import math
Register_A     = 0
Register_B     = 0x01
Register_mode  = 0x02
X_axis_H    = 0x03
Z_axis_H    = 0x05
Y_axis_H    = 0x07
eclination = -0.00669
pi          = 3.14159265359
def Magnetometer_Init():
        bus.write_byte_data(Device_Address, Register_A, 0x70)
        bus.write_byte_data(Device_Address, Register_B, 0xa0)
        bus.write_byte_data(Device_Address, Register_mode, 0)
def read_raw_data(addr):
        high = bus.read_byte_data(Device_Address, addr)
        low = bus.read_byte_data(Device_Address, addr+1)
        value = ((high << 8) | low)
        if(value > 32768):
            value = value - 65536
        return value
bus = smbus.SMBus(1) 
Device_Address = 0x1e
Magnetometer_Init() 
print (" Reading Heading Angle")
while True:
        x = read_raw_data(X_axis_H)
        z = read_raw_data(Z_axis_H)
        y = read_raw_data(Y_axis_H)
        heading = math.atan2(y, x) + declination
        if(heading > 2*pi):
                heading = heading - 2*pi
        if(heading < 0):
                heading = heading + 2*pi
        heading_angle = int(heading * 180/pi)
        print ("Heading Angle = %d°" %heading_angle)
        sleep(1)
