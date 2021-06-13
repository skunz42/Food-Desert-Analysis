import json
import urllib
import urllib.request
import sys
    
ESRI_KEY = ""
KM = 0.015

def get_resp(city, state):
    url = ('https://geocode.arcgis.com/arcgis/rest/services/World/GeocodeServer/findAddressCandidates'
        '?singleLine=%s,%s'
        '&f=%s'
        '&token=%s') % (city, state, 'pjson', ESRI_KEY)
    response = urllib.request.urlopen(url)
    jsonRaw = response.read()
    jsonData = json.loads(jsonRaw)
    return jsonData

def calc_coords(json_resp, coords):
    lngsw = json_resp["candidates"][0]["extent"]["xmin"]
    latsw = json_resp["candidates"][0]["extent"]["ymin"]
    lngne = json_resp["candidates"][0]["extent"]["xmax"]
    latne = json_resp["candidates"][0]["extent"]["ymax"]

    templng = lngsw
    templat = latsw

    while templat <= latne:
        while templng <= lngne:
            coords.append((templat, templng))
            templng += KM
        templng = lngsw
        templat += KM


def main():
    if len(sys.argv) != 3:
        print("Enter in the format: python <city> <state>")
        sys.exit(1)

    city = str(sys.argv[1])
    state = str(sys.argv[2])
    ESRI_KEY = open("esrikey.txt", "r").read()

    coords = []

    json_resp = get_resp(city, state)
    calc_coords(json_resp, coords)
    print(coords)

main()
