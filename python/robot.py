from direction import Direction


class Robot:
    x = 0
    y = 0
    direction = Direction.N

    def __init__(self, x=0, y=0, direction='N'):
        self.x = x
        self.y = y
        self.direction = Direction[direction]

    def get_current_state(self):
        return (self.x, self.y, self.direction.name)

    def rotate_clockwise(self):
        newDirection = (self.direction.value + 1) % 4
        self.direction = Direction(newDirection)
        return self.get_current_state()

    def rotate_anticlockwise(self):
        newDirection = self.direction.value - 1 if self.direction.value > 0 else 3
        self.direction = Direction(newDirection)
        return self.get_current_state()

    def move_forward(self):
        if self.direction == Direction.N:
            self.y += 1
        elif self.direction == Direction.S:
            self.y -= 1
        elif self.direction == Direction.E:
            self.x += 1
        elif self.direction == Direction.W:
            self.x -= 1
        else:
            raise Exception("Invalid Direction")
        return self.get_current_state()

