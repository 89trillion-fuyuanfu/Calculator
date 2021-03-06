from locust import HttpUser, task


class WebsiteUser(HttpUser):
    min_wait = 3000  # 执行事务之间用户等待时间的下界（单位：毫秒）
    max_wait = 6000  # 执行事务之间用户等待时间的上界（单位：毫秒

    @task
    def caculate(self):
        self.client.post("/calculate",data={'name':'2+2*6 +4/3 - 5 + 6'})
