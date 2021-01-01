struct Point {
    fn Point(x, y) {
        self.x = x
        self.y = y
    }

    fn sum {
        return self.x + self.y
    }
}

struct Vector {
    fn Vector(name, position) {
        self.name = name
        self.position = position
    }

    fn sum {
        return self.position.sum()
    }
}

vec1 = Vector{"vector", Point{12, 7}}
print(vec1.sum())
