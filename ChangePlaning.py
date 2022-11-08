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
src_info_date = ((list_work_day[0]).strftime('%d.%m.%y'))+start_time
multi_date = ((list_work_day[1]).strftime('%d.%m.%y'))+start_time
delete_vm_date = ((list_work_day[2]).strftime('%d.%m.%y'))+start_time
create_vm_date = ((list_work_day[3]).strftime('%d.%m.%y'))+start_time
setup_os_win_date = ((list_work_day[4]).strftime('%d.%m.%y'))+start_time
setup_os_nx_date = ((list_work_day[5]).strftime('%d.%m.%y'))+start_time
db_monitoring_date = ((list_work_day[6]).strftime('%d.%m.%y'))+start_time
setup_src_date = ((list_work_day[7]).strftime('%d.%m.%y'))+start_time



driver = webdriver.Firefox()
driver.get(f"https://{l}:{p}@itsm.x5.ru/sm/index.do")
time.sleep(25)


management_changes_button = driver.find_elements(By.CLASS_NAME, 'x-panel-header-text')
management_changes_button[3].click()
time.sleep(5)
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




ZNR_SRC_info = driver.find_element(By.XPATH, '//*[@id="X176_2"]')
ZNR_SRC_info.click()
time.sleep(10)
input_date_ZNR_field = driver.find_element(By.XPATH, '//*[@id="X26"]')
input_date_ZNR_field.clear()
input_date_ZNR_field.send_keys(src_info_date)
actions = ActionChains(driver)\
    .key_down(Keys.CONTROL)\
    .key_down(Keys.SHIFT)\
    .send_keys(Keys.F2)\
    .key_up(Keys.CONTROL)\
    .key_up(Keys.SHIFT)\
    .perform()

time.sleep(10)

ZNR_backup_off = driver.find_element(By.XPATH, '//*[@id="X176_3"]')
ZNR_backup_off.click()
time.sleep(10)
input_date_ZNR_field = driver.find_element(By.XPATH, '//*[@id="X26"]')
input_date_ZNR_field.clear()
input_date_ZNR_field.send_keys(multi_date)
actions = ActionChains(driver)\
    .key_down(Keys.CONTROL)\
    .key_down(Keys.SHIFT)\
    .send_keys(Keys.F2)\
    .key_up(Keys.CONTROL)\
    .key_up(Keys.SHIFT)\
    .perform()

time.sleep(10)


ZNR_bd_monitor_off = driver.find_element(By.XPATH, '//*[@id="X176_4"]')
ZNR_bd_monitor_off.click()
time.sleep(10)
input_date_ZNR_field = driver.find_element(By.XPATH, '//*[@id="X26"]')
input_date_ZNR_field.clear()
input_date_ZNR_field.send_keys(multi_date)
actions = ActionChains(driver)\
    .key_down(Keys.CONTROL)\
    .key_down(Keys.SHIFT)\
    .send_keys(Keys.F2)\
    .key_up(Keys.CONTROL)\
    .key_up(Keys.SHIFT)\
    .perform()

time.sleep(10)

ZNR_sert_delete = driver.find_element(By.XPATH, '//*[@id="X176_5"]')

actions = ActionChains(driver)\
    .send_keys(Keys.TAB)\
    .send_keys(Keys.TAB)\
    .send_keys(Keys.TAB)\
    .send_keys(Keys.TAB)\
    .send_keys(Keys.TAB)\
    .send_keys(Keys.TAB)\
    .send_keys(Keys.TAB)\
    .send_keys(Keys.TAB)\
    .send_keys(Keys.TAB)\
    .send_keys(Keys.TAB)\
    .send_keys(Keys.ENTER)\
    .perform()

time.sleep(10)

input_date_ZNR_field = driver.find_element(By.XPATH, '//*[@id="X26"]')
input_date_ZNR_field.clear()
input_date_ZNR_field.send_keys(multi_date)
actions = ActionChains(driver)\
    .key_down(Keys.CONTROL)\
    .key_down(Keys.SHIFT)\
    .send_keys(Keys.F2)\
    .key_up(Keys.CONTROL)\
    .key_up(Keys.SHIFT)\
    .perform()

time.sleep(10)


ZNR_VM_delete = driver.find_element(By.XPATH, '//*[@id="X176_6"]')
ZNR_VM_delete.click()
time.sleep(10)
input_date_ZNR_field = driver.find_element(By.XPATH, '//*[@id="X26"]')
input_date_ZNR_field.clear()
input_date_ZNR_field.send_keys(delete_vm_date)
actions = ActionChains(driver)\
    .key_down(Keys.CONTROL)\
    .key_down(Keys.SHIFT)\
    .send_keys(Keys.F2)\
    .key_up(Keys.CONTROL)\
    .key_up(Keys.SHIFT)\
    .perform()

time.sleep(10)


ZNR_VM_create = driver.find_element(By.XPATH, '//*[@id="X176_7"]')
ZNR_VM_create.click()
time.sleep(10)
input_date_ZNR_field = driver.find_element(By.XPATH, '//*[@id="X26"]')
input_date_ZNR_field.clear()
input_date_ZNR_field.send_keys(create_vm_date)
actions = ActionChains(driver)\
    .key_down(Keys.CONTROL)\
    .key_down(Keys.SHIFT)\
    .send_keys(Keys.F2)\
    .key_up(Keys.CONTROL)\
    .key_up(Keys.SHIFT)\
    .perform()

time.sleep(10)


ZNR_OS_win_setup = driver.find_element(By.XPATH, '//*[@id="X176_8"]')
ZNR_OS_win_setup.click()
time.sleep(10)
input_date_ZNR_field = driver.find_element(By.XPATH, '//*[@id="X26"]')
input_date_ZNR_field.clear()
input_date_ZNR_field.send_keys(setup_os_win_date)
actions = ActionChains(driver)\
    .key_down(Keys.CONTROL)\
    .key_down(Keys.SHIFT)\
    .send_keys(Keys.F2)\
    .key_up(Keys.CONTROL)\
    .key_up(Keys.SHIFT)\
    .perform()

time.sleep(10)


ZNR_OS_linux_setup = driver.find_element(By.XPATH, '//*[@id="X176_9"]')
ZNR_OS_linux_setup.click()
time.sleep(10)
input_date_ZNR_field = driver.find_element(By.XPATH, '//*[@id="X26"]')
input_date_ZNR_field.clear()
input_date_ZNR_field.send_keys(setup_os_nx_date)
actions = ActionChains(driver)\
    .key_down(Keys.CONTROL)\
    .key_down(Keys.SHIFT)\
    .send_keys(Keys.F2)\
    .key_up(Keys.CONTROL)\
    .key_up(Keys.SHIFT)\
    .perform()

time.sleep(10)

ZNR_bd_monitor_on = driver.find_element(By.XPATH, '//*[@id="X176_10"]')
ZNR_bd_monitor_on.click()
time.sleep(10)
input_date_ZNR_field = driver.find_element(By.XPATH, '//*[@id="X26"]')
input_date_ZNR_field.clear()
input_date_ZNR_field.send_keys(db_monitoring_date)
actions = ActionChains(driver)\
    .key_down(Keys.CONTROL)\
    .key_down(Keys.SHIFT)\
    .send_keys(Keys.F2)\
    .key_up(Keys.CONTROL)\
    .key_up(Keys.SHIFT)\
    .perform()

time.sleep(10)


ZNR_SRC_setup = driver.find_element(By.XPATH, '//*[@id="X176_11"]')
ZNR_SRC_setup.click()
time.sleep(10)
input_date_ZNR_field = driver.find_element(By.XPATH, '//*[@id="X26"]')
input_date_ZNR_field.clear()
input_date_ZNR_field.send_keys(setup_src_date)
actions = ActionChains(driver)\
    .key_down(Keys.CONTROL)\
    .key_down(Keys.SHIFT)\
    .send_keys(Keys.F2)\
    .key_up(Keys.CONTROL)\
    .key_up(Keys.SHIFT)\
    .perform()

print('Работы успешно распланированы')


# -----------------------------------------------------------------

# page_sourse = driver.page_source
# f = open("C:\Git\MyProject\Temp.txt", "w")
# f.write(page_sourse)
# f.close()

# n = 0
# for element in managament_search_changes_buton:
#     print('элемент с индексом: ', n)
#     print(element.text)
#     n += 1

# management_changes_button = driver.find_element_by_xpath('//*[@id="ext-gen-top159"]')
# managament_search_changes_buton = driver.find_element_by_xpath('/html/body/div[2]/div/div[2]/div[1]/div/div/div[4]/div[2]/div/ul/div/li[7]/div/a/span')
# start_search_button = driver.find_element_by_xpath('//*[@id="ext-gen-top2459"]') #start_search_button.click()
# save_and_exit_button = driver.find_element_by_xpath('//*[@id="ext-gen-top650"]')

# plan_tab = driver.find_element_by_xpath('//*[@id="X167_t"]')

# input_change_number_field = driver.find_element_by_xpath('//*[@id="X11"]') #input_change_number_field.send_keys(change_number) 
# input_executor_ZNR_field = driver.find_element_by_xpath('//*[@id="X12"]') #input_executor_ZNR_field.send_keys(coordinator)
# input_date_ZNR_field = driver.find_element_by_xpath('//*[@id="X26"]') #input_date_ZNR_field.send_keys(plan_date)

# ZNR_plan = driver.find_element_by_xpath('//*[@id="X176_1"]')
# ZNR_SRC_info = driver.find_element_by_xpath('//*[@id="X176_2"]')
# ZNR_backup_off = driver.find_element_by_xpath('//*[@id="X176_3"]')
# ZNR_bd_monitor_off = driver.find_element_by_xpath('//*[@id="X176_4"]')
# ZNR_sert_delete = driver.find_element_by_xpath('//*[@id="X176_5"]') 
# ZNR_VM_delete = driver.find_element_by_xpath('//*[@id="X176_6"]')
# ZNR_VM_create = driver.find_element_by_xpath('//*[@id="X176_7"]')
# ZNR_OS_win_setup = driver.find_element_by_xpath('//*[@id="X176_8"]')
# ZNR_OS_linux_setup = driver.find_element_by_xpath('//*[@id="X176_9"]')
# ZNR_bd_monitor_on = driver.find_element_by_xpath('//*[@id="X176_10"]')
# ZNR_SRC_setup = driver.find_element_by_xpath('//*[@id="X176_11"]')