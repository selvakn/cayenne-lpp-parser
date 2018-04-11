package internal

type Entry struct {
	Channel  uint8       `json:"channel"`
	DataType string      `json:"dataType"`
	Value    interface{} `json:"value"`
}

func NewEntry(channel uint8, dataType string, value interface{}) *Entry {
	return &Entry{Value: value, DataType: dataType, Channel: channel}
}

type Target struct {
	Entries []Entry `json:"entries"`
}

func (t *Target) add(channel uint8, dataType string, value interface{}) {
	t.Entries = append(t.Entries, *NewEntry(channel, dataType, value))
}

func NewTarget() *Target {
	return &Target{Entries: make([]Entry, 0)}
}

func (t *Target) Port(channel uint8, value float32) {
	t.add(channel, "Port", value)
}

func (t *Target) DigitalInput(channel, value uint8) {
	t.add(channel, "DigitalInput", value)
}

func (t *Target) DigitalOutput(channel, value uint8) {
	t.add(channel, "DigitalOutput", value)
}

func (t *Target) AnalogInput(channel uint8, value float32) {
	t.add(channel, "AnalogOutput", value)
}

func (t *Target) AnalogOutput(channel uint8, value float32) {
	t.add(channel, "AnalogOutput", value)
}

func (t *Target) Luminosity(channel uint8, value uint16) {
	t.add(channel, "AnalogOutput", value)
}

func (t *Target) Presence(channel, value uint8) {
	t.add(channel, "Presence", value)
}

func (t *Target) Temperature(channel uint8, celcius float32) {
	t.add(channel, "Temperature", celcius)
}

func (t *Target) RelativeHumidity(channel uint8, rh float32) {
	t.add(channel, "RelativeHumidity", rh)
}

func (t *Target) Accelerometer(channel uint8, x, y, z float32) {
	t.add(channel, "Accelerometer", []float32{x, y, z})
}

func (t *Target) BarometricPressure(channel uint8, hpa float32) {
	t.add(channel, "BarometricPressure", hpa)
}

func (t *Target) Gyrometer(channel uint8, x, y, z float32) {
	t.add(channel, "Gyrometer", []float32{x, y, z})
}

func (t *Target) GPS(channel uint8, latitude, longitude, altitude float32) {
	t.add(channel, "GPS", []float32{latitude, longitude, altitude})
}

func (t *Target) Voltage(channel uint8, value float32) {
	t.add(channel, "Voltage", value)
}

func (t *Target) Current(channel uint8, value float32) {
	t.add(channel, "Current", value)
}

func (t *Target) Frequency(channel uint8, value float32) {
	t.add(channel, "Frequency", value)
}

func (t *Target) Energy(channel uint8, value float32) {
	t.add(channel, "Energy", value)
}
