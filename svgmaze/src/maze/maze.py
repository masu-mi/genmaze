"""
maze
"""

from typing import Iterator,Tuple
import svgwrite
import itertools

class Maze(object):
    def __init__(self, fields):
        self.fields = fields.split("\n")
        if len(self.fields) < 3 or len(self.fields[0]) < 3:
            raise ValueError("field is too small")

    def __repr__(self):
        return self.fiedls.join("\n")

    def renderMaze(self, name: str,
                   block_size: int = 20,
                   padding: int = 20) -> svgwrite.Drawing:

        h_blocks = (len(self.fields[0])+1)/2
        v_blocks = (len(self.fields)+1)/2
        size = (
            h_blocks*block_size + padding*2,
            v_blocks*block_size + padding*2
        )

        dwg = svgwrite.Drawing(name, size=size, debug=True)
        wall = dwg.add(dwg.g(id='wall', stroke='black'))
        for w in itertools.chain(self.horizonWalls(), self.verticalWalls()):
            s, e = w[0], w[1]
            wall.add(self._line(dwg, block_size, padding, s, e))
        dwg.save()

    def _line(self, dwg, b_size, padding, start, end):
        return dwg.line(
            start=(start[0]*b_size/2 + padding, start[1]*b_size/2 + padding),
            end=(end[0]*b_size/2 + padding, end[1]*b_size/2 + padding)
        )

    def horizonWalls(self) -> Iterator[Tuple[Tuple[int, int], Tuple[int, int]]]:
        header = self.fields[0]
        xl, yl = len(header), len(self.fields)
        for i in range(0, yl, 2): # for y
            last = 0
            for j in range(1, xl-1, 2):
                if self.fields[i][j] != '#':
                    if last != j-1:
                        yield ((last, i), (j-1, i))
                    last = j+1
            yield ((last, i), (xl-1, i))

    def verticalWalls(self) -> Iterator[Tuple[Tuple[int, int], Tuple[int, int]]]:
        header = self.fields[0]
        xl, yl = len(header), len(self.fields)
        for i in range(0, xl, 2): # for x
            last = 0
            for j in range(1, yl-1, 2):
                is_start = i == 0 and j == 1
                if is_start or self.fields[j][i] != '#':
                    if last != j-1:
                        yield ((i, last), (i, j-1))
                    last = j+1
            is_goal = i == xl-1
            if is_goal:
                yield ((i, last), (i, yl-4))
            else:
                yield ((i, last), (i, yl-2))
