import requests
from requests.auth import HTTPBasicAuth


# Geoserver parameters
user = 'admin'
pwd = 'geoserver'
gs_url = 'http://localhost:8080/geoserver/rest'
workspace = 'test'

headers_zip = {'Content-Type': 'application/zip'}
headers_xml = {'Content-Type': 'application/xml; charset=UTF-8'}

# The name of the image mosaic that we will create
layer = 'sst-python'

# 1) Create the image mosaic
fileobj = open('init.zip', 'rb')

r = requests.put('{0}/workspaces/{1}/coveragestores/{2}/file.imagemosaic'\
     .format(gs_url, workspace, layer),
                 auth=HTTPBasicAuth(user, pwd),
                 files={"archive": ("init.zip", fileobj)},
                 headers=headers_zip
                 )
print(r)

# 2) Enable the time dimension
data_xml = '<coverage>\
            <enabled>true</enabled>\
            <metadata><entry key="time">\
            <dimensionInfo>\
            <enabled>true</enabled>\
            <presentation>LIST</presentation>\
            <units>ISO8601</units><defaultValue/>\
            </dimensionInfo>\
            </entry></metadata>\
            </coverage>'

r = requests.put('{0}/workspaces/{1}/coveragestores/{2}/coverages/{2}'\
     .format(gs_url, workspace, layer),
                 auth=HTTPBasicAuth(user, pwd),
                 data=data_xml,
                 headers=headers_xml
                 )
print(r)

# 3) Add new granules to the image mosaic
fileobj = open('granules.zip', 'rb')

r = requests.post('{0}/workspaces/{1}/coveragestores/{2}/file.imagemosaic?recalculate= nativebbox,latlonbbox'\
     .format(gs_url, workspace, layer),
                 auth=HTTPBasicAuth(user, pwd),
                 files={"archive": ("tiff.zip", fileobj)},
                 headers=headers_zip
                 )
print(r)
