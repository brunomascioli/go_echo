import socket
import time
from locust import User, task, between

class TCPSocketUser(User):
    wait_time = between(0.5, 2.0)
    
    def on_start(self):
        try:
            host = self.environment.host
            port = 8080 

            self.client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
            self.client.connect((host, port))
            
            self.writer = self.client.makefile('wb')
            self.reader = self.client.makefile('rb')
            
        except Exception as e:
            print(f"Falha ao conectar ao host {host}:{port}: {e}")
            self.environment.runner.quit() 

    def on_stop(self):
        if self.client:
            self.client.close()
        if self.writer:
            self.writer.close()
        if self.reader:
            self.reader.close()

    @task
    def send_and_receive_echo(self):
        request_name = "tcp_echo"
        
        message_to_send = b"Hello, Go Server!\n"
        
        start_time = time.time()
        
        try:
            self.writer.write(message_to_send)
            self.writer.flush() 
            
            response_data = self.reader.readline()
            
            total_time = int((time.time() - start_time) * 1000) 
            
            if response_data.startswith(b"Echo:"):
                self.environment.events.request.fire(
                    request_type="tcp",
                    name=request_name,
                    response_time=total_time,
                    response_length=len(response_data),
                    exception=None,
                )
            else:
                raise Exception("Resposta invalida do servidor: " + response_data.decode())

        except Exception as e:
            total_time = int((time.time() - start_time) * 1000)
            self.environment.events.request.fire(
                request_type="tcp",
                name=request_name,
                response_time=total_time,
                response_length=0,
                exception=e,
            )