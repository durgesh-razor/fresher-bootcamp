<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\TodosController;

Route::get('/todo', [TodosController::class, 'getAllTodos']);
Route::patch('/todo/{id}', [TodosController::class, 'updateTodo']);
Route::get('/todo/{id}', [TodosController::class, 'getTodo']);
Route::post('/todo', [TodosController::class, 'createTodo']);
Route::delete('/todo/{id}', [TodosController::class, 'delete']);