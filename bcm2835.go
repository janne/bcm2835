// Package bcm2835 provides functions for the bcm2835 as used in the Raspberry Pi
package bcm2835

// #include "bcm2835.h"
import "C"
import "errors"
import "unsafe"

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

// Init initialise the library by opening /dev/mem and getting pointers to the
// internal memory for BCM 2835 device registers. You must call this
// (successfully) before calling any other functions in this library (except
// SetDebug)
func Init() (err error) {
  if C.bcm2835_init() == 0 {
    return errors.New("Init: failed")
  }
  return
}

// Close closes the library, deallocating any allocaterd memory and closing
// /dev/mem
func Close() (err error) {
  if C.bcm2835_close() == 0 {
    return errors.New("Close: failed")
  }
  return
}

// SetDebug sets the debug level of the library.  A value of 1 prevents mapping
// to /dev/mem, and makes the library print out what it would do, rather than
// accessing the GPIO registers.  A value of 0, the default, causes normal
// operation.  Call this before calling Init()
func SetDebug(debug int) {
  C.bcm2835_set_debug(C.uint8_t(debug))
}

// GpioFsel sets the function select register for the given pin, which
// configures the pin as either Input or Output
func GpioFsel(pin, mode int) {
  C.bcm2835_gpio_fsel(C.uint8_t(pin), C.uint8_t(mode))
}

// GpioSet sets the specified pin output to high.
func GpioSet(pin int) {
  C.bcm2835_gpio_set(C.uint8_t(pin))
}

// GpioClr sets the specified pin output to low.
func GpioClr(pin int) {
  C.bcm2835_gpio_clr(C.uint8_t(pin))
}

// GpioLev reads the current level on the specified pin and returns either high
// or low. Works whether or not the pin is an input or an output.
func GpioLev(pin int) int {
  return int(C.bcm2835_gpio_lev(C.uint8_t(pin)))
}

func GpioEds(pin int) int {
  return int(C.bcm2835_gpio_eds(C.uint8_t(pin)))
}

func GpioSetEds(pin int) {
  C.bcm2835_gpio_set_eds(C.uint8_t(pin))
}

func GpioRen(pin int) {
  C.bcm2835_gpio_ren(C.uint8_t(pin))
}

func GpioClrRen(pin int) {
  C.bcm2835_gpio_clr_ren(C.uint8_t(pin))
}

func GpioFen(pin int) {
  C.bcm2835_gpio_fen(C.uint8_t(pin))
}

func GpioClrFen(pin int) {
  C.bcm2835_gpio_clr_fen(C.uint8_t(pin))
}

func GpioHen(pin int) {
  C.bcm2835_gpio_hen(C.uint8_t(pin))
}

func GpioClrHen(pin int) {
  C.bcm2835_gpio_clr_hen(C.uint8_t(pin))
}

func GpioLen(pin int) {
  C.bcm2835_gpio_len(C.uint8_t(pin))
}

func GpioClrLen(pin int) {
  C.bcm2835_gpio_clr_len(C.uint8_t(pin))
}

func GpioAren(pin int) {
  C.bcm2835_gpio_aren(C.uint8_t(pin))
}

func GpioClrAren(pin int) {
  C.bcm2835_gpio_clr_aren(C.uint8_t(pin))
}

func GpioAfen(pin int) {
  C.bcm2835_gpio_afen(C.uint8_t(pin))
}

func GpioClrAfen(pin int) {
  C.bcm2835_gpio_clr_afen(C.uint8_t(pin))
}

func GpioPud(pud int) {
  C.bcm2835_gpio_pud(C.uint8_t(pud))
}

func GpioPudclk(pin, on int) {
  C.bcm2835_gpio_pudclk(C.uint8_t(pin), C.uint8_t(on))
}

func GpioPad(group int) uint32 {
  return uint32(C.bcm2835_gpio_pad(C.uint8_t(group)))
}

func GpioSetPad(group int, control uint32) {
  C.bcm2835_gpio_set_pad(C.uint8_t(group), C.uint32_t(control))
}

/// GpioWrite sets the output state of the specified pin
func GpioWrite(pin, on int) {
  C.bcm2835_gpio_write(C.uint8_t(pin), C.uint8_t(on))
}

func GpioSetPud(pin, pud int) {
  C.bcm2835_gpio_set_pud(C.uint8_t(pin), C.uint8_t(pud))
}

func SpiBegin() {
  C.bcm2835_spi_begin()
}

func SpiEnd() {
  C.bcm2835_spi_end()
}

func SpiSetBitOrder(order int) {
  C.bcm2835_spi_setBitOrder(C.uint8_t(order))
}

func SpiSetClockDivider(divider uint16) {
  C.bcm2835_spi_setClockDivider(C.uint16_t(divider))
}

func SpiSetDataMode(mode int) {
  C.bcm2835_spi_setDataMode(C.uint8_t(mode))
}

func SpiChipSelect(cs int) {
  C.bcm2835_spi_chipSelect(C.uint8_t(cs))
}

func SpiSetChipSelectPolarity(cs, active int) {
  C.bcm2835_spi_setChipSelectPolarity(C.uint8_t(cs), C.uint8_t(active))
}

func SpiTransfer(value int) int {
  return int(C.bcm2835_spi_transfer(C.uint8_t(value)))
}

func SpiTransfern(data []byte) {
  C.bcm2835_spi_transfern((*C.char)(unsafe.Pointer(&data[0])), C.uint32_t(len(data)))
}

func SpiTransfernb(dst []byte, data []byte) {
  C.bcm2835_spi_transfernb((*C.char)(unsafe.Pointer(&data[0])), (*C.char)(unsafe.Pointer(&dst[0])), C.uint32_t(len(data)))
}
