#add needed libraries for a random text generation
import random
#add needed libraries for a file creation and writing
from os import path, makedirs
from sys import exit, argv
#add needed libraries for a map generation
import numpy as np


#using numeric values, because numpy array uses int8
class Tile:
    building_a = 1
    building_b = 2
    building_c = 3
    tree = 4
    alley_horiz = 5
    alley_vert = 6
    horiz_st = 7
    vert_st = 8
    police_st = 9
    hosp = 10
    bank = 11
    traphouse = 12
    gunshop = 13
    loan_shark = 14
    grass = 15
    water = 16
    dirt = 17
    metro = 18

# create a dictionary to map the integer values to their corresponding symbols
tile_symbols = {
    Tile.building_a: 'a',
    Tile.building_b: 'b',
    Tile.building_c: 'c',
    Tile.tree: '^',
    Tile.alley_horiz: '_',
    Tile.alley_vert: '!',
    Tile.horiz_st: '-',
    Tile.vert_st: '|',
    Tile.police_st: 'p',
    Tile.hosp: 'h',
    Tile.bank: '$',
    Tile.traphouse: 't',
    Tile.gunshop: 'g',
    Tile.loan_shark: 'l',
    Tile.grass: ' ',
    Tile.water: '~',
    Tile.dirt: '#',
    Tile.metro: 'm',
}

# keeping that for later in case imperative mapgen fails
#width = int(argv[1])
#length = int(argv[2])
#filename = argv[3]
class Borough:
    def __init__(self, name, width, length):
        self.name = name
        self.width = width
        self.length = length
        #self.map = np.zeros((self.length, self.width), dtype=np.int8)
        #self.map.fill(Tile.grass)

bronx = Borough('bronx', 300, 300)
queens = Borough('queens', int(bronx.width * 2.3), int(bronx.length * 2.3))
manhattan = Borough('manhattan', int(bronx.width * 1.5), int(bronx.length * 1.9))
staten_island = Borough('staten_island', int(bronx.width * 1.2), int(bronx.length * 1.2))
brooklyn = Borough('brooklyn', int(bronx.width * 1.8), int(bronx.length * 1.8))

boroughs = [bronx, queens, manhattan, staten_island, brooklyn]

for borough in boroughs:
    map_array = np.random.randint(1,19, size=(borough.length, borough.width))

def write_map(width, length, filename):
  # create a directory for the map files if it does not exist
  if not path.exists('../maps'):
    makedirs('../maps')

  # convert the integer values in the map_array to their corresponding symbols
  symbol_map_array = np.vectorize(tile_symbols.get)(map_array)
    # place the streets and alleys on the map
  x = random.randint(0, borough.width - 1)
  y = random.randint(0, borough.length - 1)
  map_array[y][x] = random.choice([Tile.horiz_st, Tile.vert_st, Tile.alley_horiz, Tile.alley_vert])
    #keep moving in a random direction until we hit the edge of the map
  while x > 0 and x < borough.width - 1 and y > 0 and y < borough.length - 1:
        dx, dy = random.choice([(0, 1), (0, -1), (1, 0), (-1, 0)])
        x += dx
        y += dy
        map_array[y][x] = random.choice([Tile.horiz_st, Tile.vert_st, Tile.alley_horiz, Tile.alley_vert])
    #place the trees on the map. Trees must be surrounded by grass. In Manhattan, we'll create Central Park
  for i in range(0, borough.width):
        for j in range(0, borough.length):
            if map_array[i][j] == Tile.grass:
                if i < 0 or i >= len(map_array):
                    continue
                if j < 0 or j >= len(map_array[i]):
                    continue
                if random.randint(0, 100) < 5:
                    map_array[i][j] = Tile.tree



  # calculate the number of tiles for each feature
  num_banks = int(4 * (width * length / (300 * 300)))
  num_hospitals = int(1 * (width * length / (300 * 300)))
  num_police_stations = int(1 * (width * length / (300 * 300)))
  num_traphouses = int(3 * (width * length / (300 * 300)))
  num_gunshops = int(2 * (width * length / (300 * 300)))
  num_metro_stations = int(1 * (width * length / (300 * 300)))
  num_parks = int(0.15 * width * length)
  num_water = int(0.1 * width * length)

  # place the features on the map
  for i in range(num_banks):
    x = random.randint(0, width-1)
    y = random.randint(0, length-1)
    symbol_map_array[y][x] = Tile.bank
  for i in range(num_hospitals):
    x = random.randint(0, width-1)
    y = random.randint(0, length-1)
    symbol_map_array[y][x] = Tile.hosp
  for i in range(num_police_stations):
    x = random.randint(0, width-1)
    y = random.randint(0, length-1)
    symbol_map_array[y][x] = Tile.police_st
  for i in range(num_traphouses):
    x = random.randint(0, width-1)
    y = random.randint(0, length-1)
    symbol_map_array[y][x] = Tile.traphouse
  for i in range(num_gunshops):
    x = random.randint(0, width-1)
    y = random.randint(0, length-1)
    symbol_map_array[y][x] = Tile.gunshop
  for i in range(num_metro_stations):
    x = random.randint(0, width-1)
    y = random.randint(0, length-1)
    symbol_map_array[y][x] = Tile.metro
  for i in range(num_parks):
    x = random.randint(0, width-1)
    y = random.randint(0, length-1)
    symbol_map_array[y][x] = Tile.grass
  for i in range(num_water):
    x = random.randint(0, width-1)
    y = random.randint(0, length-1)
    symbol_map_array[y][x] = Tile.water
  # write the map array to a file
  np.savetxt('../maps/' + filename + '.map', symbol_map_array, fmt='%s', delimiter='')

write_map(bronx.width, bronx.length, 'bronx')
write_map(manhattan.width, manhattan.length, 'manhattan')
write_map(staten.island.width, staten.island.length, 'staten.island')
write_map(brooklyn.width, brooklyn.length, 'brooklyn')
write_map(queens.width, queens.length, 'queens')
