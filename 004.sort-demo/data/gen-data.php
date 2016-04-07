<?php
$i = 0;
while ($i < 100000) {
    echo rand(1, 10000000) . chr(10);
    $i ++;
}