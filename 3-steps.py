import datetime
import os
import time
from calendar import week
from ast import Continue
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.action_chains import ActionChains
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import Select


l = (os.environ['zzz'])
p = (os.environ['zz'])


holidays = [datetime.date(2022, 6, 13), datetime.date(2022, 11, 4)]
start_time = ' 09:00:00'
coordinator = f'Витюховский Игорь ({l})'
success = 'Успешно'

change_number = input('Введите номер ИЗМа:   ')



list_work_day = []
n = 1
while list_work_day.__len__() < 8:
    date = datetime.date.today() + datetime.timedelta(days=n)
    weekday = date.weekday()
    if weekday < 5:
        if date in holidays:
            n = n+1
        else:
            list_work_day.append(date)
            n = n+1
    else:
        n = n+1

plan_date = datetime.datetime.today().strftime("%d.%m.%y")+start_time
change_policy_date = ((list_work_day[0]).strftime('%d.%m.%y'))+start_time
cancel_change_policy_date = ((list_work_day[1]).strftime('%d.%m.%y'))+start_time

driver = webdriver.Firefox()
driver.get(f"https://{l}:{p}@itsm.x5.ru/sm/index.do")
time.sleep(15)


management_changes_button = driver.find_elements(By.CLASS_NAME, 'x-panel-header-text')
management_changes_button[3].click()

managament_search_changes_buton = driver.find_elements(By.CLASS_NAME, 'x-tree-node-anchor')
managament_search_changes_buton[18].click()
time.sleep(10)

input_change_number_field = driver.switch_to.active_element
input_change_number_field.send_keys(change_number)
input_change_number_field.send_keys(Keys.ENTER)
time.sleep(10)

iFrames = driver.find_elements(By.TAG_NAME, 'iframe')
driver.switch_to.frame(iFrames[1])

plan_tab = driver.find_element(By.XPATH, '//*[@id="X167_t"]')
plan_tab.click()

ZNR_plan = driver.find_element(By.XPATH, '//*[@id="X176_1"]')
ZNR_plan.click()
time.sleep(10)


input_executor_ZNR_field = driver.find_element(By.ID, 'X12') 
input_executor_ZNR_field.clear()
time.sleep(10)
input_executor_ZNR_field.send_keys(coordinator)
time.sleep(10)
input_executor_ZNR_field.send_keys(Keys.ENTER)
time.sleep(10)

input_date_ZNR_field = driver.find_element(By.XPATH, '//*[@id="X26"]')
input_date_ZNR_field.clear()
input_date_ZNR_field.send_keys(plan_date)


input_comment_field = driver.find_element(By.XPATH, '//*[@id="X55View"]')
input_comment_field.click()

actions = ActionChains(driver)\
    .key_down(Keys.CONTROL)\
    .send_keys('a')\
    .key_up(Keys.CONTROL)\
    .send_keys(Keys.BACK_SPACE)\
    .send_keys(success)\
    .key_down(Keys.CONTROL)\
    .key_down(Keys.SHIFT)\
    .send_keys(Keys.F2)\
    .key_up(Keys.CONTROL)\
    .key_up(Keys.SHIFT)\
    .perform()

time.sleep(20)




ZNR_change_policy = driver.find_element(By.XPATH, '//*[@id="X176_2"]')
ZNR_change_policy.click()
time.sleep(10)
input_date_ZNR_field = driver.find_element(By.XPATH, '//*[@id="X26"]')
input_date_ZNR_field.clear()
input_date_ZNR_field.send_keys(change_policy_date)
actions = ActionChains(driver)\
    .key_down(Keys.CONTROL)\
    .key_down(Keys.SHIFT)\
    .send_keys(Keys.F2)\
    .key_up(Keys.CONTROL)\
    .key_up(Keys.SHIFT)\
    .perform()

time.sleep(10)

ZNR_cancel_change_policy = driver.find_element(By.XPATH, '//*[@id="X176_3"]')
ZNR_cancel_change_policy.click()
time.sleep(10)
input_date_ZNR_field = driver.find_element(By.XPATH, '//*[@id="X26"]')
input_date_ZNR_field.clear()
input_date_ZNR_field.send_keys(cancel_change_policy_date)
actions = ActionChains(driver)\
    .key_down(Keys.CONTROL)\
    .key_down(Keys.SHIFT)\
    .send_keys(Keys.F2)\
    .key_up(Keys.CONTROL)\
    .key_up(Keys.SHIFT)\
    .perform()

time.sleep(10)



print('Работы успешно распланированы')


