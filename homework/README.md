# Cara Pengerjaan Homework

```diff
- DISCLAIMER: JIKA PANDUAN INI TIDAK DIIKUTI DENGAN BAIK, MAKA HOMEWORK YANG DISUBMIT TIDAK DAPAT DINILAI.
- OTOMATIS NILAI YANG DIDAPATKAN ADALAH NOL (0).
- HATI - HATI! PANDUAN INI MERUPAKAN PANDUAN YANG HUKUMNYA WAJIB UNTUK DIIKUTI.
```

```diff
+ Silakan bertanya kepada tutor kelas masing - masing terkait panduan pengerjaan homework ini.
```

## Ketentuan Struktur Folder

Struktur folder di bawah ini adalah struktur originalnya.
Struktur ini **TIDAK BOLEH** diubah dan diganti namanya.
Struktur ini juga **TIDAK BOLEH** ditambahkan dan dikurangi.
Penggantian nama hanya diperbolehkan untuk folder `homework` yang merupakan *parent* dari **homework question folder**.

```
.
├── homework                       # your homework folder
│   ├── changoroutine              # homework question folder
│   ├── httphandler                # homework question folder
│   └── oop                        # homework question folder
└── ...
```

## Ketentuan Penamaan Folder `homework`

Penamaan folder `homework` untuk pengumpulan juga memiliki panduan khusus. Format penamaannya adalah sebagai berikut:

`KELAS_{{KELAS}}_{{NAMA}}`

Contoh:

`KELAS_C_HAFIZ_PUTRA_LUDYANTO`

`KELAS` hanya diisi **satu huruf abjadnya**, seperti contoh di atas.

`NAMA` juga memiliki ketentuan khusus, yaitu:
- **TIDAK BOLEH** mengandung karakter, selain `_`, `-`, `A-Z`, dan `a-z`. Representasi dalam regex adalah sebagai berikut: `[A-Za-z_-]`.
- Hanya boleh menggunakan huruf kapital, seperti contoh di atas.
- Nama **TIDAK BOLEH** di awali dan diakhiri dengan karakter selain `A-Z` dan `a-z`. Representasi dalam regex adalah sebagai berikut: `^[A-Za-z]+[A-Za-z-_]*[A-Za-z]+$`.
- Spasi '` `' pada nama diganti dengan ***underscore*** '`_`'.

Dengan menggunakan ketentuan penamaan di atas, struktur folder akan berubah menjadi seperti di bawah ini:

```
.
├── KELAS_C_HAFIZ_PUTRA_LUDYANTO   # your homework folder
│   ├── changoroutine              # homework question folder
│   ├── httphandler                # homework question folder
│   └── oop                        # homework question folder
└── ...
```

## Ketentuan Kode Program

Kode program yang dikembangkan untuk mengerjakan *homework* ini juga memiliki beberapa ketentuan. Ketentuan - ketentuan tersebut adalah sebagai berikut:
- Kode program hanya boleh menggunakan *library* **standar** dan **eksternal**, tetapi kode program **TIDAK BOLEH** menggunakan *library* **custom**. Contoh yang **TIDAK BOLEH** digunakan adalah sebagai berikut:
```
.
├── KELAS_C_HAFIZ_PUTRA_LUDYANTO   # your homework folder
│   ├── changoroutine              # homework question folder
│   ├── httphandler                # homework question folder
│   ├── oop                        # homework question folder
│   └── customlib                  # your custom library folder
└── ...
```
- Kode program hanya boleh ditambahkan pada **homework question folder** saja, tetapi kode program **TIDAK BOLEH** ditambahkan di luar *folder* - *folder* tersebut. Contoh yang **TIDAK BOLEH** digunakan adalah sebagai berikut:
```
.
├── KELAS_C_HAFIZ_PUTRA_LUDYANTO   # your homework folder
│   ├── changoroutine              # homework question folder
│   ├── httphandler                # homework question folder
│   ├── oop                        # homework question folder
│   └── code.go                    # your other golang code
└── ...
```