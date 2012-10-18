package bcm2835

// #include "bcm2835.h"
import "C"

const (
  Low = 0
  High = 1
  Input = 0
  Output = 1
  Pin3 = 0
  Pin5 = 1
  Pin7 = 4
  Pin8 = 14
  Pin10 = 15
  Pin11 = 17
  Pin12 = 18
  Pin13 = 21
  Pin15 = 22
  Pin16 = 23
  Pin18 = 24
  Pin19 = 10
  Pin21 = 9
  Pin22 = 25
  Pin23 = 11
  Pin24 = 8
  Pin26 = 7
)

func Init() {
  C.bcm2835_init()

  /*
  if (C.bcm2835_init() != 1) {
    return
  }
  */
}

func SetDebug(debug uint8) {
  C.bcm2835_set_debug(C.uint8_t(debug))
}

func Close() {
  C.bcm2835_close()
}

func GpioFsel(pin, mode uint8) {
  C.bcm2835_gpio_fsel(C.uint8_t(pin), C.uint8_t(mode))
}

func Delay(millis uint) {
  C.delay(C.uint(millis))
}

func GpioSet(pin uint8) {
  C.bcm2835_gpio_set(C.uint8_t(pin))
}

func GpioClr(pin uint8) {
  C.bcm2835_gpio_clr(C.uint8_t(pin))
}

func GpioLev(pin uint8) uint8 {
  return uint8(C.bcm2835_gpio_lev(C.uint8_t(pin)))
}

func GpioEds(pin uint8) uint8 {
  return uint8(C.bcm2835_gpio_eds(C.uint8_t(pin)))
}

func GpioSetEds(pin uint8) {
  C.bcm2835_gpio_set_eds(C.uint8_t(pin))
}

func GpioRen(pin uint8) {
  C.bcm2835_gpio_ren(C.uint8_t(pin))
}

func GpioClrRen(pin uint8) {
  C.bcm2835_gpio_clr_ren(C.uint8_t(pin))
}

func GpioFen(pin uint8) {
  C.bcm2835_gpio_fen(C.uint8_t(pin))
}

func GpioClrFen(pin uint8) {
  C.bcm2835_gpio_clr_fen(C.uint8_t(pin))
}

func GpioHen(pin uint8) {
  C.bcm2835_gpio_hen(C.uint8_t(pin))
}

func GpioClrHen(pin uint8) {
  C.bcm2835_gpio_clr_hen(C.uint8_t(pin))
}

func GpioLen(pin uint8) {
  C.bcm2835_gpio_len(C.uint8_t(pin))
}

func GpioClrLen(pin uint8) {
  C.bcm2835_gpio_clr_len(C.uint8_t(pin))
}

func GpioAren(pin uint8) {
  C.bcm2835_gpio_aren(C.uint8_t(pin))
}

func GpioClrAren(pin uint8) {
  C.bcm2835_gpio_clr_aren(C.uint8_t(pin))
}

func GpioAfen(pin uint8) {
  C.bcm2835_gpio_afen(C.uint8_t(pin))
}

func GpioClrAfen(pin uint8) {
  C.bcm2835_gpio_clr_afen(C.uint8_t(pin))
}

func GpioPud(pud uint8) {
  C.bcm2835_gpio_pud(C.uint8_t(pud))
}

func GpioPudclk(pin, on uint8) {
  C.bcm2835_gpio_pudclk(C.uint8_t(pin), C.uint8_t(on))
}

func GpioPad(group uint8) uint32 {
  return uint32(C.bcm2835_gpio_pad(C.uint8_t(group)))
}

func GpioSetPad(group uint8, control uint32) {
  C.bcm2835_gpio_set_pad(C.uint8_t(group), C.uint32_t(control))
}

func DelayMicroseconds(micros uint) {
  C.delayMicroseconds (C.uint(micros))
}

func GpioWrite(pin, on uint8) {
  C.bcm2835_gpio_write(C.uint8_t(pin), C.uint8_t(on))
}

func GpioSetPud(pin, pud uint8) {
  C.bcm2835_gpio_set_pud(C.uint8_t(pin), C.uint8_t(pud))
}

func SpiBegin() {
  C.bcm2835_spi_begin()
}

func SpiEnd() {
  C.bcm2835_spi_end()
}

func SpiSetBitOrder(order uint8) {
  C.bcm2835_spi_setBitOrder(C.uint8_t(order))
}

func SpiSetClockDivider(divider uint16) {
  C.bcm2835_spi_setClockDivider(C.uint16_t(divider))
}

func SpiSetDataMode(mode uint8) {
  C.bcm2835_spi_setDataMode(C.uint8_t(mode))
}

func SpiChipSelect(cs uint8) {
  C.bcm2835_spi_chipSelect(C.uint8_t(cs))
}

func SpiSetChipSelectPolarity(cs, active uint8) {
  C.bcm2835_spi_setChipSelectPolarity(C.uint8_t(cs), C.uint8_t(active))
}

func SpiTransfer(value uint8) uint8 {
  return uint8(C.bcm2835_spi_transfer(C.uint8_t(value)))
}

/*
func SpiTransfernb() {
  C.bcm2835_spi_transfernb(char* tbuf, char* rbuf, uint32_t len)
}

func SpiTransfern() {
  C.bcm2835_spi_transfern(char* buf, uint32_t len)
}
*/
