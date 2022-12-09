package modes

import (
	"image"
	"reflect"
	"testing"
)

func TestNegative(t *testing.T) {
	type args struct {
		img image.Image
	}
	tests := []struct {
		name string
		args args
		want image.Image
	}{
		{
			name: "Imagen con un solo pixel rojo",
			args: args{
				img: image.NewRGBA(image.Rect(0, 0, 1, 1)),
			},
			want: image.NewRGBA(image.Rect(0, 0, 1, 1)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Negative(tt.args.img); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Negative() = %v, want %v", got, tt.want)
			}
		})
	}
}
