const asyncHandler = require("express-async-handler");
const { getAllStudents, addNewStudent, getStudentDetail, setStudentStatus, updateStudent } = require("./students-service");

const handleGetAllStudents = asyncHandler(async (req, res) => {
    //write your code
    try{
        const student=await getAllStudents({});
        res.json({student})
    }catch{
        console.error("Error in get all student")
    }

});

const handleAddStudent = asyncHandler(async (req, res) => {
    //write your code

});

const handleUpdateStudent = asyncHandler(async (req, res) => {
    //write your code
});

const handleGetStudentDetail = asyncHandler(async (req, res) => {
    //write your code

    try{
        const student=await getStudentDetail(req.params.id);
        res.json({student});
    }catch(err){
        console.error("Error in get student details",err);
        throw err;
    }
});

const handleStudentStatus = asyncHandler(async (req, res) => {
    //write your code

});

module.exports = {
    handleGetAllStudents,
    handleGetStudentDetail,
    handleAddStudent,
    handleStudentStatus,
    handleUpdateStudent,
};
