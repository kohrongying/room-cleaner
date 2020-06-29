import unittest

from robot import Robot
from room import Room


class RoomTest(unittest.TestCase):

    def setUp(self) -> None:
        self.room = Room(6, 6, Robot())

    def test_get_robot_position(self):
        self.assertEqual((0, 0, 'N'), self.room.get_robot_position())

    def test_invalid_command(self):
        with self.assertRaises(Exception) as context:
            self.room.moveRobot('A')
            self.assertTrue('Invalid command' in str(context.exception))

    def test_move_robot_R(self):
        self.room.moveRobot('R')
        self.assertEqual((0, 0, 'E'), self.room.get_robot_position())

    def test_move_robot_L(self):
        self.room.moveRobot('L')
        self.assertEqual((0, 0, 'W'), self.room.get_robot_position())

    def test_move_robot_M(self):
        self.room.moveRobot('M')
        self.assertEqual((0, 1, 'N'), self.room.get_robot_position())

    def test_check_wall(self):
        self.assertEqual(True, self.room.is_facing_wall((0, 0, 'W')))

    def test_check_wall2(self):
        self.assertEqual(False, self.room.is_facing_wall((1, 1, 'W')))

    def test_room(self):
        room = Room(6, 6, Robot(1, 2, 'N'))
        room.setInstructions( ['L', 'M', 'L', 'M', 'L', 'M', 'L', 'M', 'M'])
        self.assertEqual((1, 3, 'N'), room.get_robot_position())

    def test_room2(self):
        room = Room(6, 6, Robot(3, 5, 'N'))
        room.setInstructions( ['M', 'L', 'M'])
        self.assertEqual((2, 5, 'W'), room.get_robot_position())


if __name__ == '__main__':
    unittest.main()
