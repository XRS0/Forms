#!/bin/sh

# Запускаем auth-service в фоне
/auth-service &

# Запускаем jwt-service в фоне
/jwt-service &

# Ждем завершения обоих процессов
wait
