package database

// Sensor представляє сенсор
type Sensor struct {
	Type            SensorType
	Range           float64
	UpdateFrequency int
}

func (s Sensor) String() string {
	return s.Model()
}

func (s Sensor) Format() string {
	if s.Model() == `T 001` {
		return `markdown`
	}

	return `html`
}

func (s Sensor) Data() string {
	switch s.Model() {
	case `T 001`:
		fallthrough
	case `T 002`:
		fallthrough
	case `T 003`:
		return `<p>Температура вранці: <strong>15°C</strong></p>
<p>Температура вдень: <strong>22°C</strong></p>
<p>Температура ввечері: <strong>18°C</strong></p>`
	case `P 2701`:
		fallthrough
	case `P 2703`:
		return `<p>Поточна температура: <strong>21°C</strong></p>
<p>Мінімальна температура: <i>10°C</i>, Максимальна температура: <i>25°C</i></p>
<p>Атмосферний тиск: <strong>1015 hPa</strong></p>`
	}

	return ``
}

func (s Sensor) Model() string {
	switch s.Type.Name {
	case `Temperature`:
		switch {
		case s.Range < 50:
			return `T 001`
		case s.Range > 50 && s.Range < 90:
			return `T 002`
		case s.Range > 91:
			return `T 003`
		}
	case `Pressure`:
		switch {
		case s.UpdateFrequency < 0:
			return `P 2701`
		case s.UpdateFrequency > 0:
			return `P 2703`
		}
	}

	return `U 00`
}

type SensorType struct {
	Name string
}

// SensorBuilder інтерфейс для побудови сенсора
type SensorBuilder interface {
	SetType(SensorType) SensorBuilder
	SetRange(float64) SensorBuilder
	SetUpdateFrequency(int) SensorBuilder
	Build() Sensor
}

// TemperatureSensorBuilder для температурних сенсорів
type SpecificSensorBuilder struct {
	sensor Sensor
}

func (b *SpecificSensorBuilder) SetType(t SensorType) SensorBuilder {
	b.sensor.Type = t
	return b
}

func (b *SpecificSensorBuilder) SetRange(r float64) SensorBuilder {
	b.sensor.Range = r
	return b
}

func (b *SpecificSensorBuilder) SetUpdateFrequency(f int) SensorBuilder {
	b.sensor.UpdateFrequency = f
	return b
}

func (b *SpecificSensorBuilder) Build() Sensor {
	return b.sensor
}

// Director відповідає за побудову сенсора
type Director struct {
	builder SensorBuilder
}

func NewDirector(b SensorBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) Construct(t SensorType, r float64, uf int) Sensor {
	return d.builder.
		SetType(t).
		SetRange(r).
		SetUpdateFrequency(uf).
		Build()
}
