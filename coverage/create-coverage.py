import requests
from requests.auth import HTTPBasicAuth


# Geoserver parameters
user = 'admin'
pwd = 'geoserver'
gs_url = 'http://localhost:8080/geoserver/rest'
workspace = 'test'
style = 'dem'
coverage = 'dem-python'

headers_tiff = {'Content-Type': 'image/tiff'}
headers_xml = {'Content-Type': 'application/xml; charset=UTF-8'}

fileobj = open('dem.tiff', 'rb')

# 1) Create the coverage
# Will create a file dem-python.geotiff in data/test/dem-python/
r = requests.put('{0}/workspaces/{1}/coveragestores/{2}/file.geotiff'.format(gs_url, workspace, coverage),
                 auth=HTTPBasicAuth(user, pwd),
                 data=fileobj,
                 headers=headers_tiff
                 )
print(r)
print(r.text)
print('\n')

# 2) Change the style
payload = '<layer><defaultStyle><name>{0}</name></defaultStyle></layer>'.format(
    style)

r = requests.put('{0}/layers/{1}'.format(gs_url, coverage),
                 auth=HTTPBasicAuth(user, pwd),
                 data=payload,
                 headers=headers_xml
                 )
print(r)
print(r.text)