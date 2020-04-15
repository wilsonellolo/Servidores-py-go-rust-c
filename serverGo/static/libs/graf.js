<%grafica.forEach(punto => {%>
    { tiempo: '<%=punto.tiempo%>', uso: <%=punto.porcentaje%> },
<%});%>


-		ab = audit_log_start(NULL, GFP_KERNEL, AUDIT_LOGIN);
-		if (ab) {
-			audit_log_format(ab, "login pid=%d" 
                "uid=%u "
-				"old auid=%u new auid=%u"
-				" old ses=%u new ses=%u",
-				task->pid,
-				from_kuid(&init_user_ns, task_uid(task)),
-				from_kuid(&init_user_ns, task->loginuid),
-				from_kuid(&init_user_ns, loginuid),
-				task->sessionid, sessionid);
-			audit_log_end(ab);
-		}

Id: 1 Nombre: systemd Estado: 1 prio: 120  prioridad en tiempo real: 0
cpu: 2 tgid:: 1 loginuid:task->loginuid sessionid: -1 ptrace 0  login pid=1 uid=0


printk("state name %s",get_state_name(task->state)); //it will print the init state first














/***************************************/

seq_printf(m,"\n Nombre: %s,&nbsp&nbsp Usuario: %u,&nbsp&nbsp Estado id: %ld,&nbsp&nbsp Nombre estado: %s ,&nbsp&nbsp Porcentaje de Ram:";
 	
seq_printf(m,"<ul class=\"nav\"><li><a><form>"); 
seq_printf(m,"<input type=\"text\" value=\"%d\" style=\"background-color: #000; border-color: #000; color: #fff;\">" +
			"<h4>Nombre: %s,&nbsp&nbsp Usuario: %u,&nbsp&nbsp Estado id: %ld,&nbsp&nbsp Nombre estado: %s ,&nbsp&nbsp Porcentaje de Ram:</h4>"+
			"<input class=\"btn btn-danger\" type=\"submit\" value=\"Matar\"></form></a>");
/***************************** */
seq_printf(m,"<ul>");					
seq_printf(m,"<li><h4>id: %d, Nombre: %s</h4></li>",task->pid,task->comm);		
seq_printf(m,"</ul>");
/************************** */
	
seq_printf(m,"</li></ul><hr>"); 