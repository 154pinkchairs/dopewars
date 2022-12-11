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


width = int(argv[1])
height = int(argv[2])
filename = argv[3]

def write_map(width, height, filename):
  # create a directory for the map files if it does not exist
  if not path.exists('../maps'):
    makedirs('../maps')

  # create a numpy array of the given dimensions
  map_array = np.zeros((height, width), dtype=np.int8)
  # convert the integer values in the map_array to their corresponding symbols
  symbol_map_array = np.vectorize(tile_symbols.get)(map_array)
  # calculate the number of tiles for each feature
  num_banks = int(4 * (width * height / (300 * 300)))
  num_hospitals = int(1 * (width * height / (300 * 300)))
  num_police_stations = int(1 * (width * height / (300 * 300)))
  num_traphouses = int(3 * (width * height / (300 * 300)))
  num_gunshops = int(2 * (width * height / (300 * 300)))
  num_metro_stations = int(1 * (width * height / (300 * 300)))
  num_parks = int(0.15 * width * height)
  num_water = int(0.1 * width * height)

  # place the features on the map
  for i in range(num_banks):
    x = random.randint(0, width-1)
    y = random.randint(0, height-1)
    symbol_map_array[y][x] = Tile.bank
  for i in range(num_hospitals):
    x = random.randint(0, width-1)
    y = random.randint(0, height-1)
    symbol_map_array[y][x] = Tile.hosp
  for i in range(num_police_stations):
    x = random.randint(0, width-1)
    y = random.randint(0, height-1)
    symbol_map_array[y][x] = Tile.police_st
  for i in range(num_traphouses):
    x = random.randint(0, width-1)
    y = random.randint(0, height-1)
    symbol_map_array[y][x] = Tile.traphouse
  for i in range(num_gunshops):
    x = random.randint(0, width-1)
    y = random.randint(0, height-1)
    symbol_map_array[y][x] = Tile.gunshop
  for i in range(num_metro_stations):
    x = random.randint(0, width-1)
    y = random.randint(0, height-1)
    symbol_map_array[y][x] = Tile.metro
  for i in range(num_parks):
    x = random.randint(0, width-1)
    y = random.randint(0, height-1)
    symbol_map_array[y][x] = Tile.grass
  for i in range(num_water):
    x = random.randint(0, width-1)
    y = random.randint(0, height-1)
    symbol_map_array[y][x] = Tile.water
  # write the map array to a file
  np.savetxt('../maps/' + filename + '.map', symbol_map_array, fmt='%s', delimiter='')

if __name__ == '__main__':
    write_map(width, height, filename)
