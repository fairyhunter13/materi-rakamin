# Cara Pengerjaan Homework

> <span style="color:red">DISCLAIMER: JIKA PANDUAN INI TIDAK DIIKUTI DENGAN BAIK, MAKA HOMEWORK YANG DISUBMIT TIDAK DAPAT DINILAI. MAKA, OTOMATIS NILAI YANG DIDAPATKAN ADALAH NOL (0). HATI - HATI! PANDUAN INI MERUPAKAN PANDUAN YANG HUKUMNYA WAJIB UNTUK DIIKUTI.</span>

> <span style="color:green">Silakan bertanya kepada tutor kelas masing - masing terkait panduan pengerjaan homework ini.</span>

## Ketentuan Struktur Folder

Struktur folder di bawah ini adalah struktur originalnya.
Struktur ini tidak boleh diubah dan diganti namanya.
Struktur ini juga tidak boleh ditambahkan dan dikurangi.
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
- **Tidak boleh** mengandung karakter, selain `_`, `-`, `A-Z`, dan `a-z`. Representasi dalam regex adalah sebagai berikut: `[A-Za-z_-]`.
- Hanya boleh menggunakan huruf kapital, seperti contoh di atas.
- Nama **tidak boleh** di awali dan diakhiri dengan karakter selain `A-Z` dan `a-z`. Representasi dalam regex adalah sebagai berikut: `^[A-Za-z]+[A-Za-z-_]*[A-Za-z]+$`.
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

### Ketentuan Kode Program