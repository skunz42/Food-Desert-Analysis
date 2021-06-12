import json
import urllib
import urllib.request
import sys

def calc_coords(key, city, state, coords):
    url = ('https://geocode.arcgis.com/arcgis/rest/services/World/GeocodeServer/findAddressCandidates'
        '?singleLine=%s,%s'
        '&f=%s'
        '&token=%s') % (city, state, 'pjson', key)
    response = urllib.request.urlopen(url)
    jsonRaw = response.read()
    jsonData = json.loads(jsonRaw)
    print(jsonData)

def main():
    if len(sys.argv) != 3:
        print("Enter in the format: python <city> <state>")
        sys.exit(1)

    city = str(sys.argv[1])
    state = str(sys.argv[2])

    ESRI_KEY = open("esrikey.txt", "r").read()
    coords = []

    calc_coords(ESRI_KEY, city, state, coords)

main()
