import unittest

from direction import Direction


class DirectionTest(unittest.TestCase):
    def test_name(self):
        self.assertEqual(Direction.N.name, 'N')

    def test_value(self):
        self.assertEqual(Direction.N.value, 0)


if __name__ == '__main__':
    unittest.main()
