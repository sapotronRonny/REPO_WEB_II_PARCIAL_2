<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     */
    public function up(): void
    {
        Schema::create('documentals', function (Blueprint $table) {
            $table->id('documental_id');
            $table->string('titulo', 255);
            $table->date('fecha_estreno')->nullable();
            $table->unsignedBigInteger('director_id')->nullable();
            $table->integer('duracion')->nullable();
            $table->text('descripcion')->nullable();
            $table->timestamps();

            $table->foreign('director_id')->references('director_id')->on('directores')->onDelete('set null');
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('documentals');
    }
};
