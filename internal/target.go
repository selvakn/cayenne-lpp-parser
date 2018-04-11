package internal

type Target struct {
	values map[uint8]interface{}
}

func NewTarget() *Target{
	return &Target{values: make(map[uint8]interface{})}
}

func (t *Target) Port(channel uint8, value float32) {
	t.values[channel] = value
}

func (t *Target) DigitalInput(channel, value uint8) {
	t.values[channel] = value
}

func (t *Target) DigitalOutput(channel, value uint8) {
	t.values[channel] = value
}

func (t *Target) AnalogInput(channel uint8, value float32) {
	t.values[channel] = value
}

func (t *Target) AnalogOutput(channel uint8, value float32) {
	t.values[channel] = value
}

func (t *Target) Luminosity(channel uint8, value uint16) {
	t.values[channel] = value
}

func (t *Target) Presence(channel, value uint8) {
	t.values[channel] = value
}

func (t *Target) Temperature(channel uint8, celcius float32) {
	t.values[channel] = celcius
}

func (t *Target) RelativeHumidity(channel uint8, rh float32) {
	t.values[channel] = rh
}

func (t *Target) Accelerometer(channel uint8, x, y, z float32) {
	t.values[channel] = []float32{x, y, z}
}

func (t *Target) BarometricPressure(channel uint8, hpa float32) {
	t.values[channel] = hpa
}

func (t *Target) Gyrometer(channel uint8, x, y, z float32) {
	t.values[channel] = []float32{x, y, z}
}

func (t *Target) GPS(channel uint8, latitude, longitude, altitude float32) {
	t.values[channel] = []float32{latitude, longitude, altitude}
}

func (t *Target) Voltage(channel uint8, value float32) {
	t.values[channel] = value
}

func (t *Target) Current(channel uint8, value float32) {
	t.values[channel] = value
}

func (t *Target) Frequency(channel uint8, value float32) {
	t.values[channel] = value
}

func (t *Target) Energy(channel uint8, value float32) {
	t.values[channel] = value
}

