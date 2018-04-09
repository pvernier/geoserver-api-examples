# Coverage

## Illustrates how to create a coverage (raster layer).

The examples contain 2 steps:

* Upload a GeoTIFF file called dem.tiff
* Change the style of the created coverage to a style called 'dem'

**NB**: It's possible to upload the GeoTIFF file inside a ZIP file.
In this case just replace the content-type (use 'application/zip' instead of 'image/tiff').

The cURL commands are below. For the Python and Go examples see respected files in this folder.

### 1) Create the coverage
```
# Will create a file dem-curl.geotiff in data/test/dem-curl/
curl -v -u 'admin:geoserver' -XPUT -H "Content-type: image/tif" --data-binary @dem.tiff http://localhost:8080/geoserver/rest/workspaces/test/coveragestores/dem-curl/file.geotiff
```

### 2) Change the style (to a style called 'dem')
```
curl -v -u 'admin:geoserver' -XPUT -H "Content-type:application/xml; charset=UTF-8" -d '<layer><defaultStyle><name>dem</name></defaultStyle></layer>' http://localhost:8080/geoserver/rest/layers/dem-curl
```

**NB**: if your Geoserver is on HTTPS, end your curl command with: -k