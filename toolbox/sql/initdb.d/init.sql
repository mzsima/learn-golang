DROP SCHEMA IF EXISTS sample;
CREATE SCHEMA sample;
USE sample;

DROP TABLE IF EXISTS candlestick;

CREATE TABLE candlestick
(
  id           INT(10),
  open    decimal(34,20),
  high    decimal(34,20),
  low     decimal(34,20),
  close   decimal(34,20),
  volume  decimal(34,20)
);

INSERT INTO candlestick (id, open, high, low, close, volume) VALUES (1, 6764919, 6764940, 6758680, 6764900, 0.2674); 
INSERT INTO candlestick (id, open, high, low, close, volume) VALUES (2, 6764940, 6764940, 6764940, 6764940, 0.0002); 
INSERT INTO candlestick (id, open, high, low, close, volume) VALUES (3, 6764920, 6764920, 6764920, 6764920, 0.04); 
INSERT INTO candlestick (id, open, high, low, close, volume) VALUES (4, 6764920, 6764920, 6764920, 6764920, 0.002); 
INSERT INTO candlestick (id, open, high, low, close, volume) VALUES (5, 6764920, 6769879, 6764920, 6769879, 0.12); 
INSERT INTO candlestick (id, open, high, low, close, volume) VALUES (6, 6760961, 6777656, 6760961, 6777656, 1.342); 
INSERT INTO candlestick (id, open, high, low, close, volume) VALUES (7, 6761441, 6772754, 6761441, 6772299, 0.0626); 
INSERT INTO candlestick (id, open, high, low, close, volume) VALUES (8, 6772257, 6772258, 6772257, 6772258, 0.008); 
INSERT INTO candlestick (id, open, high, low, close, volume) VALUES (9, 6765244, 6765244, 6765244, 6765244, 0.018); 
INSERT INTO candlestick (id, open, high, low, close, volume) VALUES (10, 6765640, 6772259, 6765640, 6765661, 0.0296); 
INSERT INTO candlestick (id, open, high, low, close, volume) VALUES (11, 6772754, 6772754, 6772754, 6772754, 0.0008); 
INSERT INTO candlestick (id, open, high, low, close, volume) VALUES (12, 6772680, 6772680, 6772680, 6772680, 0.0002); 
INSERT INTO candlestick (id, open, high, low, close, volume) VALUES (13, 6766576, 6771560, 6765009, 6765009, 0.1646); 
INSERT INTO candlestick (id, open, high, low, close, volume) VALUES (14, 6765009, 6769140, 6761460, 6761460, 0.044); 
INSERT INTO candlestick (id, open, high, low, close, volume) VALUES (15, 6767850, 6767850, 6767518, 6767518, 0.022); 
INSERT INTO candlestick (id, open, high, low, close, volume) VALUES (16, 6766320, 6766320, 6766300, 6766300, 0.004); 
