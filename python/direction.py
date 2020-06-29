from enum import Enum


class Direction(Enum):
    N = 0
    E = 1
    S = 2
    W = 3

    # def get_string(self):
    #     d = {
    #         0: 'N',
    #         1: 'E',
    #         2: 'S',
    #         3: 'W'
    #     }
    #     return d[self.value]
