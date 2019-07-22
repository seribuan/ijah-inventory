## Ijah Inventory REST API

Used third party package:
```
go get github.com/gin-gonic/gin
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/sqlite
```
Please put this application in $GOPATH/sorabel

After that, to run, simply run:
```
go run main.go
```

```
[GIN-debug] GET    /barang/:id               --> sorabel/controllers.(*BarangController).Get-fm (3 handlers)
[GIN-debug] GET    /barang                   --> sorabel/controllers.(*BarangController).GetAll-fm (3 handlers)
[GIN-debug] POST   /barang                   --> sorabel/controllers.(*BarangController).Post-fm (3 handlers)
[GIN-debug] PUT    /barang/:id               --> sorabel/controllers.(*BarangController).Put-fm (3 handlers)
[GIN-debug] GET    /barang.csv               --> sorabel/controllers.(*BarangController).Export-fm (3 handlers)
[GIN-debug] GET    /barang_masuk/:id         --> sorabel/controllers.(*BarangMasukController).Get-fm (3 handlers)
[GIN-debug] GET    /barang_masuk             --> sorabel/controllers.(*BarangMasukController).GetAll-fm (3 handlers)
[GIN-debug] POST   /barang_masuk             --> sorabel/controllers.(*BarangMasukController).Post-fm (3 handlers)
[GIN-debug] GET    /barang_masuk.csv         --> sorabel/controllers.(*BarangMasukController).Export-fm (3 handlers)
[GIN-debug] GET    /barang_keluar/:id        --> sorabel/controllers.(*BarangKeluarController).Get-fm (3 handlers)
[GIN-debug] GET    /barang_keluar            --> sorabel/controllers.(*BarangKeluarController).GetAll-fm (3 handlers)
[GIN-debug] POST   /barang_keluar            --> sorabel/controllers.(*BarangKeluarController).Post-fm (3 handlers)
[GIN-debug] GET    /barang_keluar.csv        --> sorabel/controllers.(*BarangKeluarController).Export-fm (3 handlers)
[GIN-debug] GET    /laporan/barang           --> sorabel/controllers.(*LaporanController).GetLaporanNilai-fm (3 handlers)
[GIN-debug] GET    /laporan/barang.csv       --> sorabel/controllers.(*LaporanController).GetLaporanNilaiCSV-fm (3 handlers)
[GIN-debug] GET    /laporan/penjualan        --> sorabel/controllers.(*LaporanController).GetLaporanPenjualan-fm (3 handlers)
[GIN-debug] GET    /laporan/penjualan.csv    --> sorabel/controllers.(*LaporanController).GetLaporanPenjualanCSV-fm (3 handlers)
```

This is the list of the API.
For the PUT (update) API, I didn't implement it based on the assumption it shouldn't be allowed to edit Kwitansi (but the code is there). I also didn't implement delete for the same assumption

For the data, I use internal auto-increment IDs as the primary key, so it will be easier to generate unique key

For the Laporan Penjualan, I just noticed that it has from-end date filters, that I haven't implemented.

Sample post data:
```
curl -XPOST -d 'sku=SSI-D00791077-MM-ABC&nama=Zalekia Plain Casual Blouse (M,Fine White)' 'localhost:8080/barang' for barang
curl -XPOST -d 'no_kwitansi=KW125&barang_id=4&jumlah_pesan=200&jumlah_diterima:100&catatan=20171103-98876&harga=55000' 'localhost:8080/barang_masuk'
curl -XPOST -d 'barang_id=4&jumlah_keluar=1&catatan=Hilang&harga=100000' 'localhost:8080/barang_keluar'
```

For the import, right now there is no REST API, but it can be run in main.go by running:
```
barangWriter := dbwriter.NewBarangWriter(db)
importer.ImportBarang("Toko Ijah.xlsx", barangWriter)
barangMasukWriter := dbwriter.NewBarangMasukWriter(db)
importer.ImportBarangMasuk("Toko Ijah.xlsx", barangMasukWriter)
barangKeluarWriter := dbwriter.NewBarangKeluarWriter(db)
importer.ImportBarangKeluar("Toko Ijah.xlsx", barangKeluarWriter)
```

I use the design for the Import Functionality so that it can be extended easily, for other things that's not DB
(I mistakenly thought that the import is a mandatory feature, that's why in the log it can be seen that I worked on it first)

The compiled binary is called also included, called main
