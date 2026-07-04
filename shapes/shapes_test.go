package shapes

import "testing"

// ──────────────────────────────────────────────────────────────────────────
//  Тесты. Менять их НЕ нужно — твоя задача сделать так, чтобы они прошли.
//  Это table-driven тесты: один цикл прогоняет много случаев. Так пишут
//  тесты в идиоматичном Go, и так же сделано в курсе.
// ──────────────────────────────────────────────────────────────────────────

func TestArea(t *testing.T) {
	cases := []struct {
		name string
		shape Shape
		want float64
	}{
		{name: "прямоугольник", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "круг", shape: Circle{Radius: 10}, want: 314.1592653589793},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.shape.Area()
			if got != c.want {
				t.Errorf("%s: получили %g, ожидали %g", c.name, got, c.want)
			}
		})
	}
}

func TestTotalArea(t *testing.T) {
	shapes := []Shape{
		Rectangle{Width: 12, Height: 6}, // 72
		Circle{Radius: 10},              // ~314.159
	}

	got := TotalArea(shapes)
	want := 386.1592653589793 // 72 + π*100

	if got != want {
		t.Errorf("TotalArea: получили %g, ожидали %g", got, want)
	}
}
