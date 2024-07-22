//Permets d'obliger le joueur à choisir une réponse.
function validateForm() {
    var options = document.getElementsByName("answer");
    var isChecked = false;
        
    for (var i = 0; i < options.length; i++) {
        if (options[i].checked) {
            isChecked = true;
            break;
        }
    }   
    
    if (!isChecked) {
        alert("Veuillez sélectionner une réponse !");
        return false; 
    }
}