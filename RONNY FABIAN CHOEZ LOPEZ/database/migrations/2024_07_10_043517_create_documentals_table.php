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
            $table->id('id_documental');
            $table->string('titulo', 100);
            $table->foreignId('id_genero')->constrained('genero_docu');
            $table->foreignId('id_director')->constrained('directors');
            $table->date('fecha_lanzamiento')->nullable();
            $table->integer('duracion')->nullable();  // duración en minutos
            $table->text('resumen')->nullable();
            $table->timestamps();

            $table->foreign('id_genero')->references('id_genero')->on('genero_docu');
            $table->foreign('id_director')->references('id_director')->on('directors');
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
