import unittest

from robot import Robot


class RobotTest(unittest.TestCase):

    def setUp(self):
        self.robot = Robot()

    def test_set_initial_robot_position(self):
        customRobot = Robot(1, 2, 'E')
        self.assertEqual((1, 2, 'E'), customRobot.get_current_state())

    def test_get_current_state(self):
        self.assertEqual(self.robot.get_current_state(), (0, 0, 'N'))

    def test_rotate_clockwise(self):
        self.assertEqual(self.robot.rotate_clockwise(), (0, 0, 'E'))

    def test_rotate_clockwise_at_west(self):
        self.robot.rotate_clockwise()
        self.robot.rotate_clockwise()
        self.robot.rotate_clockwise()
        self.assertEqual(self.robot.rotate_clockwise(), (0, 0, 'N'))

    def test_rotate_anticlockwise(self):
        self.assertEqual(self.robot.rotate_anticlockwise(), (0, 0, 'W'))

    def test_move_forward_when_N(self):
        self.assertEqual((0, 1, 'N'), self.robot.move_forward())

    def test_move_forward_when_E(self):
        self.robot.rotate_clockwise()
        self.assertEqual((1, 0, 'E'), self.robot.move_forward())

    def test_move_forward_when_S(self):
        self.robot.rotate_clockwise()
        self.robot.rotate_clockwise()
        self.assertEqual((0, -1, 'S'), self.robot.move_forward())

    def test_move_forward_when_W(self):
        self.robot.rotate_anticlockwise()
        self.assertEqual((-1, 0, 'W'), self.robot.move_forward())


if __name__ == '__main__':
    unittest.main()
