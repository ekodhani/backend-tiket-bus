-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 24 Sep 2023 pada 04.04
-- Versi server: 10.4.28-MariaDB
-- Versi PHP: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_abtb`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `bis`
--

CREATE TABLE `bis` (
  `kode_bis` int(11) NOT NULL,
  `nama_bis` varchar(40) NOT NULL,
  `kursi` varchar(20) NOT NULL,
  `jam_berangkat` time NOT NULL,
  `kelas` enum('Premium','Reguler') NOT NULL,
  `harga` int(25) NOT NULL,
  `keterangan` text NOT NULL,
  `rute` text NOT NULL,
  `image` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `bis`
--

INSERT INTO `bis` (`kode_bis`, `nama_bis`, `kursi`, `jam_berangkat`, `kelas`, `harga`, `keterangan`, `rute`, `image`) VALUES
(10, 'TransJabodetabek Jakarta', '20', '09:00:00', 'Premium', 30000, 'Ac, Hiburan, Kursi Recliner, Hotspot\r\n\r\nRute\r\n\r\nBundaran HI - Summarecon Mall Bekasi \r\n', 'Jakarta - Bekasi', 'TJimage.jpg'),
(11, 'TransJabodetabek Jakarta', '20', '15:00:00', 'Reguler', 15000, 'Ac, Kursi Recliner\r\n\r\nRute\r\n\r\nBundaran HI - Summarecon Mall Bekasi', 'Jakarta - Bekasi', 'TJimage.jpg'),
(12, 'TransJabodetabek Jakarta', '20', '08:00:00', 'Premium', 30000, 'Ac, Hiburan, Kursi Recliner, Hotspot', 'Jakarta - Tangerang', 'TJimage.jpg'),
(13, 'TransJabodetabek Jakarta', '20', '15:00:00', 'Reguler', 15000, 'Ac, Kursi Recliner', 'Jakarta - Tangerang', 'TJimage.jpg'),
(14, 'TransJabodetabek Jakarta', '20', '09:00:00', 'Premium', 30000, 'Ac, Hiburan, Kursi Recliner, Hotspot', 'Jakarta - Depok', 'TJimage.jpg'),
(15, 'TransJabodetabek Jakarta', '20', '16:00:00', 'Reguler', 15000, 'Ac, Kursi Recliner', 'Jakarta - Depok', 'TJimage.jpg'),
(16, 'TransJabodetabek Jakarta', '20', '07:00:00', 'Premium', 30000, 'Ac, Hiburan, Kursi Recliner, Hotspot', 'Jakarta - Bogor', 'TJimage.jpg'),
(17, 'TransJabodetabek Jakarta', '20', '12:30:00', 'Reguler', 15000, 'Ac, Kursi Recliner', 'Jakarta - Bogor', 'TJimage.jpg'),
(18, 'TransJabodetabek Bekasi', '20', '06:40:00', 'Premium', 30000, 'Ac, Hiburan, Kursi Recliner, Hotspot\r\nRute\r\nSummarecon Mall Bekasi - Bundaran HI', 'Bekasi - Jakarta', 'TJimageBksi.jpg'),
(19, 'TransJabodetabek Bekasi', '20', '05:00:00', 'Reguler', 15000, 'Ac, Kursi Recliner\r\nRute\r\nSummarecon Mall Bekasi - Bundaran HI', 'Bekasi - Jakarta', 'TJimageBksi.jpg'),
(20, 'TransJabodetabek Bekasi', '20', '07:30:00', 'Premium', 30000, 'Ac, Hiburan, Kursi Recliner, Hotspot', 'Bekasi - Tangerang', 'TJimageBksi.jpg'),
(21, 'TransJabodetabek Bekasi', '20', '09:00:00', 'Reguler', 15000, 'Ac, Kursi Recliner', 'Bekasi - Tangerang', 'TJimageBksi.jpg');

-- --------------------------------------------------------

--
-- Struktur dari tabel `bus`
--

CREATE TABLE `bus` (
  `id` int(11) NOT NULL,
  `nama_bus` varchar(100) NOT NULL,
  `nomor_polisi` varchar(50) NOT NULL,
  `jumlah_kursi` int(11) NOT NULL,
  `type` varchar(50) NOT NULL,
  `harga` int(15) NOT NULL,
  `status` int(2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `bus`
--

INSERT INTO `bus` (`id`, `nama_bus`, `nomor_polisi`, `jumlah_kursi`, `type`, `harga`, `status`) VALUES
(1, 'Sumber Jaya', 'B4B1LU', 20, 'ekonomi', 200000, 1);

-- --------------------------------------------------------

--
-- Struktur dari tabel `detail_pemesanan_tiket`
--

CREATE TABLE `detail_pemesanan_tiket` (
  `id` int(11) NOT NULL,
  `id_tiket` int(11) NOT NULL,
  `nama` varchar(100) NOT NULL,
  `nik` int(11) NOT NULL,
  `id_kursi` int(11) NOT NULL,
  `is_cancel` int(11) NOT NULL,
  `cancel_date` int(11) NOT NULL,
  `cancel_by` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `detail_pemesanan_tiket`
--

INSERT INTO `detail_pemesanan_tiket` (`id`, `id_tiket`, `nama`, `nik`, `id_kursi`, `is_cancel`, `cancel_date`, `cancel_by`) VALUES
(1, 1, 'dhani', 123131, 1, 0, 0, ''),
(2, 1, 'rere', 1231231, 4, 0, 0, ''),
(3, 1, 'haikal', 6576576, 5, 0, 0, ''),
(4, 5, 'dhani', 132132, 1, 0, 0, ''),
(5, 5, 'dhani', 23131, 4, 0, 0, '');

-- --------------------------------------------------------

--
-- Struktur dari tabel `detail_pesan`
--

CREATE TABLE `detail_pesan` (
  `id_pesan` varchar(11) NOT NULL,
  `kode_bis` int(11) NOT NULL,
  `harga` varchar(25) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `detail_pesan`
--

INSERT INTO `detail_pesan` (`id_pesan`, `kode_bis`, `harga`) VALUES
('ABTB001', 10, 'Rp.30.000'),
('ABTB002', 11, 'Rp.25.000');

-- --------------------------------------------------------

--
-- Struktur dari tabel `kursi`
--

CREATE TABLE `kursi` (
  `id` int(11) NOT NULL,
  `id_bis` int(11) NOT NULL,
  `no_kursi` int(11) NOT NULL,
  `status` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `kursi`
--

INSERT INTO `kursi` (`id`, `id_bis`, `no_kursi`, `status`) VALUES
(1, 1, 1, 0),
(2, 1, 2, 0),
(3, 1, 3, 0),
(4, 1, 4, 0),
(5, 1, 5, 0),
(6, 1, 6, 0),
(7, 1, 7, 0),
(8, 1, 8, 0),
(9, 1, 9, 0),
(10, 1, 10, 0),
(11, 1, 11, 0),
(12, 1, 12, 0),
(13, 1, 13, 0),
(14, 1, 14, 0),
(15, 1, 15, 0),
(16, 1, 16, 0),
(17, 1, 17, 0),
(18, 1, 18, 0),
(19, 1, 19, 0),
(20, 1, 20, 0);

-- --------------------------------------------------------

--
-- Struktur dari tabel `pemesanan`
--

CREATE TABLE `pemesanan` (
  `id_pesan` varchar(11) NOT NULL,
  `email` varchar(35) NOT NULL,
  `kode_bis` int(11) NOT NULL,
  `rute` text NOT NULL,
  `harga` int(255) NOT NULL,
  `tgl_berangkat` date NOT NULL,
  `jam_berangkat` time NOT NULL,
  `kursi` enum('1','2','3','4') NOT NULL,
  `status` enum('sudah terbayar','menunggu pembayaran') NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `pemesanan`
--

INSERT INTO `pemesanan` (`id_pesan`, `email`, `kode_bis`, `rute`, `harga`, `tgl_berangkat`, `jam_berangkat`, `kursi`, `status`) VALUES
('ABTB005', 'muhammadilyas@gmail.com', 11, 'Jakarta - Bekasi', 15000, '2019-06-24', '15:00:00', '2', 'menunggu pembayaran'),
('ABTB006', 'muhammadilyas@gmail.com', 14, 'Jakarta - Depok', 30000, '2019-06-30', '09:00:00', '3', 'menunggu pembayaran'),
('ABTB007', 'muhammadbiben@gmail.com', 12, 'Jakarta - Tangerang', 30000, '2019-06-24', '08:00:00', '1', 'menunggu pembayaran'),
('ABTB008', 'muhammadbiben@gmail.com', 19, 'Bekasi - Jakarta', 15000, '2019-06-30', '05:00:00', '4', 'menunggu pembayaran'),
('ABTB009', 'sofyanarifin@gmail.com', 20, 'Bekasi - Tangerang', 30000, '2019-06-24', '07:30:00', '4', 'menunggu pembayaran'),
('ABTB010', 'sofyanarifin@gmail.com', 21, 'Bekasi - Tangerang', 15000, '2019-06-30', '09:00:00', '1', 'menunggu pembayaran'),
('ABTB011', 'awalnp@gmail.com', 16, 'Jakarta - Bogor', 30000, '2019-06-24', '07:00:00', '2', 'menunggu pembayaran'),
('ABTB012', 'awalnp@gmail.com', 17, 'Jakarta - Bogor', 15000, '2019-06-30', '12:30:00', '2', 'menunggu pembayaran'),
('ABTB013', 'muhammadilyas@gmail.com', 11, 'Jakarta - Bekasi', 15000, '2019-06-17', '15:00:00', '2', 'menunggu pembayaran'),
('ABTB014', 'muhammadilyas@gmail.com', 11, 'Jakarta - Bekasi', 15000, '2019-06-18', '15:00:00', '1', 'menunggu pembayaran'),
('ABTB015', 'muhammadilyas@gmail.com', 10, 'Jakarta - Bekasi', 30000, '2019-06-20', '09:00:00', '3', 'menunggu pembayaran'),
('ABTB016', 'muhammadilyas@gmail.com', 10, 'Jakarta - Bekasi', 30000, '2019-06-27', '09:00:00', '1', 'menunggu pembayaran'),
('ABTB017', 'muhammadilyas@gmail.com', 15, 'Jakarta - Depok', 15000, '2023-01-31', '16:00:00', '2', 'menunggu pembayaran');

-- --------------------------------------------------------

--
-- Struktur dari tabel `pemesanan_tiket`
--

CREATE TABLE `pemesanan_tiket` (
  `id` int(11) NOT NULL,
  `id_bus` int(11) NOT NULL,
  `kota_asal` varchar(50) NOT NULL,
  `kota_tujuan` varchar(50) NOT NULL,
  `id_pembayaran` int(11) NOT NULL,
  `pergi` int(11) NOT NULL,
  `pulang` int(11) NOT NULL,
  `status_pembayaran` int(11) NOT NULL,
  `id_user` int(11) NOT NULL,
  `create_date` int(11) NOT NULL,
  `create_by` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `pemesanan_tiket`
--

INSERT INTO `pemesanan_tiket` (`id`, `id_bus`, `kota_asal`, `kota_tujuan`, `id_pembayaran`, `pergi`, `pulang`, `status_pembayaran`, `id_user`, `create_date`, `create_by`) VALUES
(1, 1, 'Tangerang', 'Purworejo', 1, 1690891728, 0, 0, 1, 1693310967, 'dhani'),
(2, 1, 'Tangerang', 'Purworejo', 1, 1690854795, 0, 0, 1, 1695520481, 'dhani'),
(3, 1, 'Tangerang', 'Purworejo', 1, 1690854795, 0, 0, 1, 1695520485, 'dhani'),
(4, 1, 'Tangerang', 'Purworejo', 1, 1690854795, 0, 0, 1, 1695520494, 'dhani'),
(5, 1, 'Tangerang', 'Purworejo', 1, 1690854795, 0, 0, 1, 1695520633, 'dhani');

-- --------------------------------------------------------

--
-- Struktur dari tabel `role`
--

CREATE TABLE `role` (
  `id_role` int(1) NOT NULL,
  `nama_role` varchar(15) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `role`
--

INSERT INTO `role` (`id_role`, `nama_role`) VALUES
(1, 'administrator'),
(2, 'penumpang');

-- --------------------------------------------------------

--
-- Struktur dari tabel `rute`
--

CREATE TABLE `rute` (
  `id` int(11) NOT NULL,
  `id_bus` int(11) NOT NULL,
  `asal` varchar(100) NOT NULL,
  `tujuan` varchar(100) NOT NULL,
  `tgl_berangkat` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `rute`
--

INSERT INTO `rute` (`id`, `id_bus`, `asal`, `tujuan`, `tgl_berangkat`) VALUES
(1, 1, 'Tangerang', 'Purworejo', 1691733456);

-- --------------------------------------------------------

--
-- Struktur dari tabel `transaksi`
--

CREATE TABLE `transaksi` (
  `id_pesan` varchar(11) NOT NULL,
  `email` varchar(35) NOT NULL,
  `kode_bis` int(11) NOT NULL,
  `tgl_berangkat` date NOT NULL,
  `kursi` enum('1','2','3','4') NOT NULL,
  `status` enum('belum terbayar','sudah terbayar') NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `transaksi`
--

INSERT INTO `transaksi` (`id_pesan`, `email`, `kode_bis`, `tgl_berangkat`, `kursi`, `status`) VALUES
('ABTB000', 'sofyanarifin@gmail.com', 20, '2019-05-01', '2', 'sudah terbayar'),
('ABTB001', 'muhammadilyas@gmail.com', 12, '2019-06-23', '2', 'sudah terbayar'),
('ABTB002', 'awal.np@gmail.com', 16, '2019-06-01', '3', 'sudah terbayar'),
('ABTB003', 'muhammadbiben@gmail.com', 10, '2019-06-05', '2', 'sudah terbayar'),
('ABTB004', 'muhammadilyas@gmail.com', 17, '2019-06-01', '2', 'sudah terbayar');

-- --------------------------------------------------------

--
-- Struktur dari tabel `user`
--

CREATE TABLE `user` (
  `id_user` int(2) NOT NULL,
  `nama` varchar(35) NOT NULL,
  `no_telp` varchar(20) NOT NULL,
  `email` varchar(40) NOT NULL,
  `password` varchar(50) NOT NULL,
  `id_role` int(1) NOT NULL,
  `image` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `user`
--

INSERT INTO `user` (`id_user`, `nama`, `no_telp`, `email`, `password`, `id_role`, `image`) VALUES
(1, 'dhani', '089521649714', 'ekonurramadhani@gmail.com', '1ce7557b2f3f12fbaedfa6bcb809d662', 1, 'default.jpg');

-- --------------------------------------------------------

--
-- Struktur dari tabel `user_token`
--

CREATE TABLE `user_token` (
  `id` int(11) NOT NULL,
  `email` varchar(100) NOT NULL,
  `token` varchar(255) NOT NULL,
  `date_created` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `bis`
--
ALTER TABLE `bis`
  ADD PRIMARY KEY (`kode_bis`);

--
-- Indeks untuk tabel `bus`
--
ALTER TABLE `bus`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `detail_pemesanan_tiket`
--
ALTER TABLE `detail_pemesanan_tiket`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `kursi`
--
ALTER TABLE `kursi`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `pemesanan`
--
ALTER TABLE `pemesanan`
  ADD PRIMARY KEY (`id_pesan`);

--
-- Indeks untuk tabel `pemesanan_tiket`
--
ALTER TABLE `pemesanan_tiket`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `role`
--
ALTER TABLE `role`
  ADD PRIMARY KEY (`id_role`);

--
-- Indeks untuk tabel `rute`
--
ALTER TABLE `rute`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id_user`),
  ADD KEY `id_role` (`id_role`);

--
-- Indeks untuk tabel `user_token`
--
ALTER TABLE `user_token`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `bis`
--
ALTER TABLE `bis`
  MODIFY `kode_bis` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;

--
-- AUTO_INCREMENT untuk tabel `bus`
--
ALTER TABLE `bus`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT untuk tabel `detail_pemesanan_tiket`
--
ALTER TABLE `detail_pemesanan_tiket`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT untuk tabel `kursi`
--
ALTER TABLE `kursi`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=21;

--
-- AUTO_INCREMENT untuk tabel `pemesanan_tiket`
--
ALTER TABLE `pemesanan_tiket`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT untuk tabel `role`
--
ALTER TABLE `role`
  MODIFY `id_role` int(1) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `rute`
--
ALTER TABLE `rute`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT untuk tabel `user`
--
ALTER TABLE `user`
  MODIFY `id_user` int(2) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT untuk tabel `user_token`
--
ALTER TABLE `user_token`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
