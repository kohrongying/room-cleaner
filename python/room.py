from enum import Enum

from direction import Direction
from robot import Robot


class Room:
    width = 0
    length = 0
    robot = None

    def __init__(self, width, length, robot):
        self.width = width  # y
        self.length = length  # x
        self.robot = robot

    def moveRobot(self, command):
        if command == Command.M.value:
            if self.is_facing_wall(self.robot.get_current_state()) is False:
                self.robot.move_forward()
        elif command == Command.L.value:
            self.robot.rotate_anticlockwise()
        elif command == Command.R.value:
            self.robot.rotate_clockwise()
        else:
            raise Exception("Invalid command")

    def get_robot_position(self):
        return self.robot.get_current_state()

    def is_facing_wall(self, tup):
        (x, y, direction) = tup
        if direction == Direction.N.name and y == self.width - 1:
            return True
        if direction == Direction.W.name and x == 0:
            return True
        if direction == Direction.S.name and y == 0:
            return True
        if direction == Direction.E.name and x == self.length - 1:
            return True
        return False

    def setInstructions(self, moveset):
        for move in moveset:
            self.moveRobot(move)


class Command(Enum):
    M = 'M'
    L = 'L'
    R = 'R'
