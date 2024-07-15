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
            $table->unsignedBigInteger('id_genero');
            $table->unsignedBigInteger('id_director');
            $table->integer('duracion')->nullable();
            $table->text('resumen')->nullable();
            $table->timestamps();

            $table->foreign('id_genero')->references('id_genero')->on('genero_docus');
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
