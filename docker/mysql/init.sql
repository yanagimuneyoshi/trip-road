CREATE TABLE IF NOT EXISTS plans (
  id          INT AUTO_INCREMENT PRIMARY KEY,
  title       VARCHAR(255) NOT NULL,
  description TEXT,
  prefecture  VARCHAR(50)  NOT NULL,
  days        INT          NOT NULL DEFAULT 1,
  created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- サンプルデータ
INSERT INTO plans (title, description, prefecture, days) VALUES
('京都古都めぐり', '金閣寺・清水寺・嵐山を巡る定番コース', '京都府', 2),
('沖縄リゾート満喫', '美ら海水族館と慶良間諸島シュノーケリング', '沖縄県', 3);
