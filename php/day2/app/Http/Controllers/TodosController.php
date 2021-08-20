<?php

namespace App\Http\Controllers;

use App\Models\Todos;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

class TodosController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        //
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {
        //
    }

    /**
     * Display the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function edit($id)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, $id)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        //
    }

    public function getAllTodos(): JsonResponse
    {
        $todos = DB::table('todos')->get();
        return response()->json([
            'todos'=>$todos,
        ]);
    }

    public function getTodo(int $id): JsonResponse
    {
        $todo = DB::table('todos')->find($id);
        return response()->json([
            'todo'=>$todo,
        ]);
    }

    public function createTodo(Request $request): JsonResponse
    {
        $title = \request('title');
        $desc = \request('description');

        $data = array('title'=>$title,'description'=>$desc);
        $id = DB::table('todos')->insertGetId($data);
        return response()->json([
            'id'=>$id,
            'first_name'=>$title,
            'last_name'=>$desc]);
    }

    public function delete(int $id): JsonResponse
    {
        DB::table('todos')->delete($id);
        return response()->json([
            'message'=>"TODO deleted successfully",
        ]);
    }

    public function updateTodo(Request $request, int $id): JsonResponse
    {
        $title = \request('title');
        $desc = \request('description');

        $todo = DB::table('todos')->where('id', $id);
        if($title) {
            $todo->update(['title' => $title]);
        }
        if($desc){
            $todo->update(['description' => $desc]);
        }

        // $todo->save();

        return response()->json([
            'message'=>"TODO Updated successfully",
        ]);
    		
    }

}