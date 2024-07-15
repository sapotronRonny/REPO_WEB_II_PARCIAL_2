<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\GeneroDocuController;
use App\Http\Controllers\DirectorController;
use App\Http\Controllers\DocumentalController;

Route::apiResource('genero_docus', GeneroDocuController::class);
Route::apiResource('directors', DirectorController::class);
Route::apiResource('documentals', DocumentalController::class);
