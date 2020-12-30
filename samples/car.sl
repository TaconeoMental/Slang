// Basic program to show the use of structs

struct Car {
    fn Car(s, m) {
        self.speed = s
        self.max = m
    }
    
    fn accelerate(valor) {
        if valor + self.speed > self.max {
            return "Too fast!"
        }
        self.speed += valor
        return self.speed
    }

    fn stop {
        self.speed = 0
    }
}

new_car = Car{10, 30}

print new_car.accelerate(10)
print new_car.accelerate(10)

new_car.stop()
print new_car.speed
