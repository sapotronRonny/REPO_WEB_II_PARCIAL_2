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
        Schema::create('documental_generos', function (Blueprint $table) {
            $table->unsignedBigInteger('documental_id');
            $table->unsignedBigInteger('generodocu_id');
            $table->timestamps();

            $table->primary(['documental_id', 'generodocu_id']);
            $table->foreign('documental_id')->references('documental_id')->on('documentales')->onDelete('cascade');
            $table->foreign('generodocu_id')->references('generdocuo_id')->on('generos')->onDelete('cascade');
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('documental_generos');
    }
};
